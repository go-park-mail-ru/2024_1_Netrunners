package service

import (
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
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
