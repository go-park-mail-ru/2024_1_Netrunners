package repository

import (
	"context"

	"fmt"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PgxIface interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Close()
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type StatStorage struct {
	pool PgxIface
}

func NewStatStorage(pool PgxIface) (*StatStorage, error) {
	return &StatStorage{
		pool: pool,
	}, nil
}

const getAllProfileQuestions = `
		SELECT uuid, question FROM $1`

const getAdditionalQuestionByUuid = `
		SELECT external_id, question FROM additional_question WHERE question_external_id = $1;`

const getAllVarsByAdditionalQueUuid = `
		SELECT id_inside_question, variant FROM additional_vars WHERE question_external_id = $1;`

const putStat = `
		INSERT INTO $1 (question_external_id, score, is_additional_score) VALUES ($2, $3, $4);`

const getStatByUuid = `
		SELECT AVG(score), COUNT(score) FROM $1 WHERE question_external_id = $2 AND is_additional_score = FALSE;`

const getAddStat = `
		SELECT COUNT(score) FROM $1 WHERE question_external_id = $2 AND 
		                                  is_additional_score = TRUE AND 
		                                  id_inside_question = $3;`

const getVarById = `
 SELECT id_inside_question, variant FROM $1 WHERE id = $2 
                                              AND is_additional_score = TRUE 
                                              AND external_id = $3;`

func (storage *StatStorage) AddQuestion(question domain.AddQuestion) error {
	return nil
}

func (storage *StatStorage) GetPageQuestions(page string) ([]domain.Question, error) {
	var rows pgx.Rows
	var err error
	switch {
	case page == "profileData":
		rows, err = storage.pool.Query(context.Background(), getAllProfileQuestions, "profile_question")
	case page == "actorData":
		rows, err = storage.pool.Query(context.Background(), getAllProfileQuestions, "actor_question")
	case page == "filmsAll":
		rows, err = storage.pool.Query(context.Background(), getAllProfileQuestions, "film_question")
	case page == "filmData":
		rows, err = storage.pool.Query(context.Background(), getAllProfileQuestions, "film_data_question")
	}
	var allQuestions []domain.Question
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction to add film: %w: %w", err,
			myerrors.ErrNoSuchItemInTheCache)
	}
	for rows.Next() {
		var (
			que   domain.Question
			uuid  string
			title string
		)
		err = rows.Scan(&uuid, &title)
		que.Uuid = uuid
		que.Title = title

		rowsAddQue, err := storage.pool.Query(context.Background(), getAdditionalQuestionByUuid, que.Uuid)
		if err != nil {
			return nil, fmt.Errorf("failed to begin transaction to add film: %w: %w", err,
				myerrors.ErrNoSuchItemInTheCache)
		}
		for rowsAddQue.Next() {
			var (
				addQue     domain.AdditionalQuestion
				addQueUuid string
				addQueText string
			)
			err = rowsAddQue.Scan(&addQueUuid, &addQueText)
			if err != nil {
				return nil, fmt.Errorf("failed to begin transaction to add film: %w: %w", err,
					myerrors.ErrNoSuchItemInTheCache)
			}
			addQue.Uuid = addQueUuid
			addQue.Title = addQueText
			var variants []domain.Variant
			rowsVars, err := storage.pool.Query(context.Background(), getAllVarsByAdditionalQueUuid, addQue.Uuid)
			if err != nil {
				return nil, fmt.Errorf("failed to begin transaction to add film: %w: %w", err,
					myerrors.ErrNoSuchItemInTheCache)
			}
			for rowsVars.Next() {
				var (
					variant  domain.Variant
					id       uint32
					question string
				)
				err = rowsVars.Scan(&id, &question)
				if err != nil {
					return nil, fmt.Errorf("failed to begin transaction to add film: %w: %w", err,
						myerrors.ErrNoSuchItemInTheCache)
				}
				variant.Id = id
				variant.Title = question
				variants = append(variants, variant)
				addQue.CheckVars = variants
			}
			que.AdditionalQuestion = addQue
		}
		allQuestions = append(allQuestions, que)
	}
	return allQuestions, nil
}

func (storage *StatStorage) AddStatistics(page string, statistics []domain.AddQuestionStatistics) error {
	for _, stat := range statistics {
		var err error
		switch {
		case page == "profileData":
			_, err = storage.pool.Exec(context.Background(), putStat, "profile_question", stat.Uuid,
				stat.Score, stat.IsAdditional)
		case page == "actorData":
			_, err = storage.pool.Exec(context.Background(), putStat, "actor_question", stat.Uuid,
				stat.Score, stat.IsAdditional)
		case page == "filmsAll":
			_, err = storage.pool.Exec(context.Background(), putStat, "film_question", stat.Uuid,
				stat.Score, stat.IsAdditional)
		case page == "filmData":
			_, err = storage.pool.Exec(context.Background(), putStat, "film_data_question", stat.Uuid,
				stat.Score, stat.IsAdditional)
		}
		if err != nil {
			return fmt.Errorf("failed to add stat: %w: %w", err,
				myerrors.ErrNoSuchItemInTheCache)
		}
	}
	return nil
}

func (storage *StatStorage) GetStatisticsByPage(page string) ([]domain.QuestionStatistics, error) {
	var rows pgx.Rows
	var err error
	switch {
	case page == "profileData":
		rows, err = storage.pool.Query(context.Background(), getAllProfileQuestions, "profile_question")
	case page == "actorData":
		rows, err = storage.pool.Query(context.Background(), getAllProfileQuestions, "actor_question")
	case page == "filmsAll":
		rows, err = storage.pool.Query(context.Background(), getAllProfileQuestions, "film_question")
	case page == "filmData":
		rows, err = storage.pool.Query(context.Background(), getAllProfileQuestions, "film_data_question")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction to add film: %w: %w", err,
			myerrors.ErrNoSuchItemInTheCache)
	}
	var stats []domain.QuestionStatistics
	for rows.Next() {
		var (
			qStat     domain.QuestionStatistics
			que       domain.Question
			uuid      string
			title     string
			stat      float32
			statCount uint32
		)
		err = rows.Scan(&uuid, &title)
		que.Uuid = uuid
		switch {
		case page == "profileData":
			err = storage.pool.QueryRow(context.Background(), getStatByUuid, "profile_question",
				que.Uuid).Scan(&stat, &statCount)
		case page == "actorData":
			err = storage.pool.QueryRow(context.Background(), getStatByUuid, "actor_question",
				que.Uuid).Scan(&stat, &statCount)
		case page == "filmsAll":
			err = storage.pool.QueryRow(context.Background(), getStatByUuid, "film_question",
				que.Uuid).Scan(&stat, &statCount)
		case page == "filmData":
			err = storage.pool.QueryRow(context.Background(), getStatByUuid, "film_data_question",
				que.Uuid).Scan(&stat, &statCount)
		}
		qStat.Title = title
		qStat.IsAdditional = false
		qStat.AverageScore = stat
		qStat.ScoresCount = statCount
		stats = append(stats, qStat)
	}
	for rows.Next() {
		var (
			qStat      domain.QuestionStatistics
			que        domain.Question
			uuid       string
			title      string
			statCount  int32
			checkStat  domain.CheckQuestionStatistics
			checkStats []domain.CheckQuestionStatistics
		)
		err = rows.Scan(&uuid, &title)
		que.Uuid = uuid
		for i := 0; i < 4; i++ {
			var varTitle string
			switch {
			case page == "profileData":
				err = storage.pool.QueryRow(context.Background(), getAddStat, "profile_question",
					que.Uuid, i+1).Scan(&statCount)
				err = storage.pool.QueryRow(context.Background(), getVarById, "", i+1, que.Uuid).Scan(&varTitle)
			case page == "actorData":
				err = storage.pool.QueryRow(context.Background(), getAddStat, "actor_question",
					que.Uuid, i+1).Scan(&statCount)
				err = storage.pool.QueryRow(context.Background(), getVarById, "", i+1, que.Uuid).Scan(&varTitle)
			case page == "filmsAll":
				err = storage.pool.QueryRow(context.Background(), getAddStat, "film_question",
					que.Uuid, i+1).Scan(&statCount)
				err = storage.pool.QueryRow(context.Background(), getVarById, "", i+1, que.Uuid).Scan(&varTitle)
			case page == "filmData":
				err = storage.pool.QueryRow(context.Background(), getAddStat, "film_data_question",
					que.Uuid, i+1).Scan(&statCount)
				err = storage.pool.QueryRow(context.Background(), getVarById, "", i+1, que.Uuid).Scan(&varTitle)
			}
			checkStat.Title = varTitle
			checkStat.Count = statCount
			checkStats = append(checkStats, checkStat)
		}

		qStat.Title = title
		qStat.IsAdditional = false
		qStat.CheckVariants = checkStats
		stats = append(stats, qStat)
	}
	return stats, nil
}
