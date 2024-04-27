package repository

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
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

func (storage *StatStorage) AddQuestion(question domain.AddQuestion) error {
	return nil
}
func (storage *StatStorage) GetPageQuestions(page string) ([]domain.Question, error) {
	return []domain.Question{}, nil
}

func (storage *StatStorage) AddStatistics(statistics []domain.AddQuestionStatistics) error {
	return nil
}

func (storage *StatStorage) GetStatisticsByPage(page string) ([]domain.QuestionStatistics, error) {
	return []domain.QuestionStatistics{}, nil
}
