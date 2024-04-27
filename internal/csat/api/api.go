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
	AddStatistics(ctx context.Context, page string, statistics []domain.AddQuestionStatistics) error
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

func (csatServer *CsatServer) GetStatisticsByPage(ctx context.Context, request *session.GetStatisticsRequest) (*session.GetStatisticsResponse, error) {
	statistics, err := csatServer.csatService.GetStatisticsByPage(ctx, request.Page)
	if err != nil {
		return nil, err
	}

	return convertGetStatisticsToProto(statistics), nil
}

func (csatServer *CsatServer) GetQuestionsByPage(ctx context.Context, request *session.GetQuestionsByPageRequest) (*session.QuestionsResponse, error) {
	questions, err := csatServer.csatService.GetPageQuestions(ctx, request.Page)
	if err != nil {
		return nil, err
	}

	return convertQuestionsToProto(questions), nil
}

func (csatServer *CsatServer) AddStatistics(ctx context.Context, request *session.AddStatisticsRequest) (*session.AddStatisticsResponse, error) {
	err := csatServer.csatService.AddStatistics(ctx, request.Page, convertNewStatisticsToDomain(request.NewStatistics))
	if err != nil {
		return &session.AddStatisticsResponse{}, err
	}

	return &session.AddStatisticsResponse{}, nil
}

func convertGetStatisticsToProto(statistics []domain.QuestionStatistics) *session.GetStatisticsResponse {
	response := make([]*session.Statistics, 0, len(statistics))

	for _, stat := range statistics {
		newStat := &session.Statistics{
			Title:        stat.Title,
			IsAdditional: stat.IsAdditional,
			ScoresCount:  stat.ScoresCount,
			AverageScore: stat.AverageScore,
		}

		newVariants := make([]*session.CheckQuestionStatistics, 0, len(stat.CheckVariants))
		for _, variant := range stat.CheckVariants {
			newVariants = append(newVariants, &session.CheckQuestionStatistics{
				Title: variant.Title,
				Count: variant.Count,
			})
		}

		newStat.CheckStatistics = newVariants

		response = append(response, newStat)
	}

	return &session.GetStatisticsResponse{
		Statistics: response,
	}
}

func convertQuestionsToProto(questions []domain.Question) *session.QuestionsResponse {
	response := make([]*session.Question, 0, len(questions))

	for _, question := range questions {
		newQuestion := &session.Question{
			Uuid:  question.Uuid,
			Title: question.Title,
			AdditionalQuestion: &session.AdditionalQuestion{
				Uuid:  question.AdditionalQuestion.Uuid,
				Title: question.AdditionalQuestion.Title,
			},
		}

		newVariants := make([]*session.Variant, 0, len(question.AdditionalQuestion.CheckVars))
		for _, variant := range question.AdditionalQuestion.CheckVars {
			newVariants = append(newVariants, &session.Variant{
				Id:    variant.Id,
				Title: variant.Title,
			})
		}

		newQuestion.AdditionalQuestion.CheckVars = newVariants

		response = append(response, newQuestion)
	}

	return &session.QuestionsResponse{
		Questions: response,
	}
}

func convertNewStatisticsToDomain(statistics []*session.NewStatistics) []domain.AddQuestionStatistics {
	newStatistics := make([]domain.AddQuestionStatistics, 0, len(statistics))

	for _, stat := range statistics {
		newStatistics = append(newStatistics, domain.AddQuestionStatistics{
			Uuid:         stat.Uuid,
			IsAdditional: stat.IsAdditional,
			Score:        stat.Score,
		})
	}

	return newStatistics
}
