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
		SELECT uuid, question FROM profile_question`

const getAdditionalQuestionByUuid = `
		SELECT external_id, question FROM additional_question WHERE question_external_id = $1;`

func (storage *StatStorage) AddQuestion(question domain.AddQuestion) error {

}

func (storage *StatStorage) GetPageQuestions(page string) ([]domain.Question, error) {
	switch {
	case page == "profileData":
		var allQuestions []domain.Question
		rows, err := storage.pool.Query(context.Background(), getAllProfileQuestions)
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

			var additionalQuestions domain.AdditionalQuestion
			err = storage.pool.QueryRow(context.Background(), getAdditionalQuestionByUuid, que.Uuid)
			for rows.Next() {
				var (
					addQue     domain.AdditionalQuestion
					addQueUuid string
					addQueText string
				)
				err = rows.Scan(&addQueUuid, &addQueText)
				addQue.Uuid = addQueUuid
				addQue.Title = addQueText
				var variants []string
			}
			allQuestions = append(allQuestions, que)
		}
	case page == "actorData":

	case page == "filmsAll":
	case page == "filmData":
	}
}

func (storage *StatStorage) AddStatistics(statistics []domain.AddQuestionStatistics) error {

}

func (storage *StatStorage) GetStatisticsByPage(page string) ([]domain.QuestionStatistics, error) {

}
