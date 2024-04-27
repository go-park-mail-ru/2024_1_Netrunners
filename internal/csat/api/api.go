package api

import (
	"context"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
	"go.uber.org/zap"
)

type CsatService interface {
	AddQuestion(ctx context.Context, question domain.AddQuestion) error
	GetPageQuestions(ctx context.Context, page string) ([]domain.Question, error)
	AddStatistics(ctx context.Context, statistics []domain.AddQuestionStatistics) error
	GetStatisticsByPage(ctx context.Context, page string) ([]domain.QuestionStatistics, error)
}

type CsatServer struct {
	csatService CsatService
	logger      *zap.SugaredLogger
}

func NewCsatServer(service CsatService, logger *zap.SugaredLogger) *CsatServer {
	return &CsatServer{
		csatService: service,
		logger:      logger,
	}
}

func (csatServer *CsatServer) GetStatistics(context.Context, *session.GetStatisticsRequest) (*session.GetStatisticsResponse, error) {
	return nil, nil
}

func (csatServer *CsatServer) GetQuestionsByPage(context.Context, *session.GetQuestionsByPageRequest) (*session.QuestionsResponse, error) {
	// return nil, nil
	return &session.QuestionsResponse{
		Questions: []*session.Question{},
	}, nil
}

func (csatServer *CsatServer) AddStatistics(context.Context, *session.AddStatisticsRequest) (*session.AddStatisticsResponse, error) {
	return nil, nil
}
