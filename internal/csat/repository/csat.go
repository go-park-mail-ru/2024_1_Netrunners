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
		SELECT external_id, question FROM profile_question;`

const getAllActorQuestions = `
		SELECT external_id, question FROM actor_question;`

const getAllFilmAllQuestions = `
		SELECT external_id, question FROM film_question;`

const getAllFilmDataQuestions = `
		SELECT external_id, question FROM film_data_question;`

const getAdditionalQuestionByUuid = `
		SELECT external_id, question FROM additional_question WHERE question_external_id = $1;`

const getAllVarsByAdditionalQueUuid = `
		SELECT id_inside_question, variant FROM additional_answer WHERE question_external_id = $1;`

const putStatProfile = `
		INSERT INTO profile_stat (question_external_id, score, is_additional_score) VALUES ($1, $2, $3);`

const putStatActor = `
		INSERT INTO actor_stat (question_external_id, score, is_additional_score) VALUES ($1, $2, $3);`

const putStatFilm = `
		INSERT INTO film_data_stat (question_external_id, score, is_additional_score) VALUES ($1, $2, $3);`

const putStatFilmsAll = `
		INSERT INTO film_stat (question_external_id, score, is_additional_score) VALUES ($1, $2, $3);`

const getStatByUuidProfile = `
		SELECT AVG(score), COUNT(score) FROM profile_stat WHERE question_external_id = $1 AND is_additional_score = FALSE;`

const getStatByUuidActor = `
		SELECT AVG(score), COUNT(score) FROM actor_stat WHERE question_external_id = $1 AND is_additional_score = FALSE;`

const getStatByUuidFilmsAll = `
		SELECT AVG(score), COUNT(score) FROM film_stat WHERE question_external_id = $1 AND is_additional_score = FALSE;`

const getStatByUuidFilm = `
		SELECT AVG(score), COUNT(score) FROM film_data_stat WHERE question_external_id = $1 AND is_additional_score = FALSE;`

func (storage *StatStorage) AddQuestion(question domain.AddQuestion) error {
	return nil
}

func (storage *StatStorage) GetPageQuestions(page string) ([]domain.Question, error) {
	switch {
	case page == "profileData":
		rows, err := storage.pool.Query(context.Background(), getAllProfileQuestions)
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
			allQuestions = append(allQuestions, que)
		}
		return allQuestions, nil
	case page == "actorData":
		rows, err := storage.pool.Query(context.Background(), getAllActorQuestions)
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

			allQuestions = append(allQuestions, que)
		}
		return allQuestions, nil
	case page == "filmsAll":
		rows, err := storage.pool.Query(context.Background(), getAllFilmAllQuestions)
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
			allQuestions = append(allQuestions, que)
		}
		return allQuestions, nil
	case page == "filmData":
		rows, err := storage.pool.Query(context.Background(), getAllFilmDataQuestions, "film_data_question")
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

			allQuestions = append(allQuestions, que)
		}
		return allQuestions, nil
	}
	return nil, fmt.Errorf("failed to find question: %w", myerrors.ErrNoSuchItemInTheCache)
}

func (storage *StatStorage) AddStatistics(page string, statistics []domain.AddQuestionStatistics) error {
	for _, stat := range statistics {
		var err error
		switch {
		case page == "profileData":
			_, err = storage.pool.Exec(context.Background(), putStatProfile, stat.Uuid,
				stat.Score, stat.IsAdditional)
			if err != nil {
				return fmt.Errorf("failed to add stat: %w: %w", err,
					myerrors.ErrNoSuchItemInTheCache)
			}
		case page == "actorData":
			_, err = storage.pool.Exec(context.Background(), putStatActor, stat.Uuid,
				stat.Score, stat.IsAdditional)
			if err != nil {
				return fmt.Errorf("failed to add stat: %w: %w", err,
					myerrors.ErrNoSuchItemInTheCache)
			}
		case page == "filmsAll":
			_, err = storage.pool.Exec(context.Background(), putStatFilmsAll, stat.Uuid,
				stat.Score, stat.IsAdditional)
			if err != nil {
				return fmt.Errorf("failed to add stat: %w: %w", err,
					myerrors.ErrNoSuchItemInTheCache)
			}
		case page == "filmData":
			_, err = storage.pool.Exec(context.Background(), putStatFilm, stat.Uuid,
				stat.Score, stat.IsAdditional)
			if err != nil {
				return fmt.Errorf("failed to add stat: %w: %w", err,
					myerrors.ErrNoSuchItemInTheCache)
			}
		}
	}
	return nil
}

func (storage *StatStorage) GetStatisticsByPage(page string) ([]domain.QuestionStatistics, error) {
	switch {
	case page == "profileData":
		rows, err := storage.pool.Query(context.Background(), getAllProfileQuestions)
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
			err = storage.pool.QueryRow(context.Background(), getStatByUuidProfile, que.Uuid).Scan(&stat, &statCount)
			qStat.Title = title
			qStat.IsAdditional = false
			qStat.AverageScore = stat
			qStat.ScoresCount = statCount
			fmt.Println(qStat)
			stats = append(stats, qStat)
		}
		return stats, nil
	case page == "actorData":
		rows, err := storage.pool.Query(context.Background(), getAllActorQuestions)
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
			err = storage.pool.QueryRow(context.Background(), getStatByUuidActor, que.Uuid).Scan(&stat, &statCount)
			qStat.Title = title
			qStat.IsAdditional = false
			qStat.AverageScore = stat
			qStat.ScoresCount = statCount
			fmt.Println(qStat)
			stats = append(stats, qStat)
		}
		return stats, nil
	case page == "filmsAll":
		rows, err := storage.pool.Query(context.Background(), getAllFilmAllQuestions)
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
			err = storage.pool.QueryRow(context.Background(), getStatByUuidFilmsAll, que.Uuid).Scan(&stat, &statCount)
			qStat.Title = title
			qStat.IsAdditional = false
			qStat.AverageScore = stat
			qStat.ScoresCount = statCount
			fmt.Println(qStat)
			stats = append(stats, qStat)
		}
		return stats, nil
	case page == "filmData":
		rows, err := storage.pool.Query(context.Background(), getAllFilmDataQuestions)
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
			err = storage.pool.QueryRow(context.Background(), getStatByUuidFilm, que.Uuid).Scan(&stat, &statCount)
			qStat.Title = title
			qStat.IsAdditional = false
			qStat.AverageScore = stat
			qStat.ScoresCount = statCount
			fmt.Println(qStat)
			stats = append(stats, qStat)
		}
		return stats, nil
	}
	return nil, myerrors.ErrNoSuchItemInTheCache
}
