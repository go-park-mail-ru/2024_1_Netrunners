package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
	"go.uber.org/zap"
)

type CsatStorage interface {
	AddQuestion(question domain.AddQuestion) error
	GetPageQuestions(page string) ([]domain.Question, error)
	AddStatistics(statistics []domain.AddQuestionStatistics) error
	GetStatisticsByPage(page string) ([]domain.QuestionStatistics, error)
}

type CsatService struct {
	storage CsatStorage
	logger  *zap.SugaredLogger
}

func NewCsatService(storage CsatStorage, logger *zap.SugaredLogger) *CsatService {
	return &CsatService{
		storage: storage,
		logger:  logger,
	}
}

func (csatService *CsatService) AddQuestion(ctx context.Context, question domain.AddQuestion) error {
	err := csatService.storage.AddQuestion(question)
	if err != nil {
		csatService.logger.Errorf("[reqid=%s] failed to add question: %w", ctx.Value(requestId.ReqIDKey), err)
		return err
	}

	return nil
}

func (csatService *CsatService) GetPageQuestions(ctx context.Context, page string) ([]domain.Question, error) {
	questions, err := csatService.storage.GetPageQuestions(page)
	if err != nil {
		csatService.logger.Errorf("[reqid=%s] failed to get page questions: %w", ctx.Value(requestId.ReqIDKey), err)
		return []domain.Question{}, err
	}

	return questions, nil
}

func (csatService *CsatService) AddStatistics(ctx context.Context, statistics []domain.AddQuestionStatistics) error {
	err := csatService.storage.AddStatistics(statistics)
	if err != nil {
		csatService.logger.Errorf("[reqid=%s] failed to add statistics: %w", ctx.Value(requestId.ReqIDKey), err)
		return err
	}

	return nil
}

func (csatService *CsatService) GetStatisticsByPage(ctx context.Context, page string) ([]domain.QuestionStatistics, error) {
	questions, err := csatService.storage.GetStatisticsByPage(page)
	if err != nil {
		csatService.logger.Errorf("[reqid=%s] failed to get statistics by page: %w", ctx.Value(requestId.ReqIDKey), err)
		return []domain.QuestionStatistics{}, err
	}

	return questions, nil
}
