package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	reqid "github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
	"go.uber.org/zap"
)

type addStatisticsRequest struct {
	Page       string                         `json:"page"`
	Statistics []domain.AddQuestionStatistics `json:"statistics"`
}

type getPageQuestionsResponse struct {
	Status    int               `json:"status"`
	Questions []domain.Question `json:"questions"`
}

type addStatisticsResponse struct {
	Status int `json:"status"`
}

type getStatisticsByPageResponse struct {
	Status     int                         `json:"status"`
	Statistics []domain.QuestionStatistics `json:"statistics"`
}

type CsatHandlers struct {
	csatClient *session.CsatClient
	logger     *zap.SugaredLogger
}

func NewCsatHandlers(csatClient *session.CsatClient, logger *zap.SugaredLogger) *CsatHandlers {
	return &CsatHandlers{
		csatClient: csatClient,
		logger:     logger,
	}
}

func (csatHandlers *CsatHandlers) GetPageQuestions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	questions, err := (*csatHandlers.csatClient).GetQuestionsByPage(ctx, convertGetPageQuestionsToProto(r))
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			csatHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	domainQuestions := make([]domain.Question, 0, len(questions.Questions))
	for _, question := range questions.Questions {
		newQuestion := domain.Question{
			Uuid:  question.Uuid,
			Title: question.Title,
			AdditionalQuestion: domain.AdditionalQuestion{
				Uuid:  question.AdditionalQuestion.Uuid,
				Title: question.AdditionalQuestion.Title,
			},
		}

		newVariants := make([]domain.Variant, 0, len(question.AdditionalQuestion.CheckVars))
		for _, variant := range question.AdditionalQuestion.CheckVars {
			newVariants = append(newVariants, domain.Variant{
				Id:    variant.Id,
				Title: variant.Title,
			})
		}

		newQuestion.AdditionalQuestion.CheckVars = newVariants

		domainQuestions = append(domainQuestions, newQuestion)
	}

	response := getPageQuestionsResponse{
		Status:    http.StatusOK,
		Questions: domainQuestions,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			csatHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			csatHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}

func (csatHandlers *CsatHandlers) AddStatistics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	protoStatistics, err := convertAddStatisticsToProto(r)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			csatHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	_, err = (*csatHandlers.csatClient).AddStatistics(ctx, protoStatistics)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			csatHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	response := addStatisticsResponse{
		Status: http.StatusOK,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			csatHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			csatHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}

func (csatHandlers *CsatHandlers) GetStatisticsByPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	statistics, err := (*csatHandlers.csatClient).GetStatisticsByPage(ctx, convertGetStatisticsByPageToProto(r))
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			csatHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	domainStatistics := make([]domain.QuestionStatistics, 0, len(statistics.Statistics))
	for _, stat := range statistics.Statistics {
		newStatistics := domain.QuestionStatistics{
			Title:        stat.Title,
			IsAdditional: stat.IsAdditional,
			ScoresCount:  stat.ScoresCount,
			AverageScore: stat.AverageScore,
		}

		newCheckStatistics := make([]domain.CheckQuestionStatistics, 0, len(stat.CheckStatistics))
		for _, checkStat := range stat.CheckStatistics {
			newCheckStatistics = append(newCheckStatistics, domain.CheckQuestionStatistics{
				Title: checkStat.Title,
				Count: checkStat.Count,
			})
		}

		newStatistics.CheckVariants = newCheckStatistics

		domainStatistics = append(domainStatistics, newStatistics)
	}

	response := getStatisticsByPageResponse{
		Status:     http.StatusOK,
		Statistics: domainStatistics,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			csatHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			csatHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}

func convertGetPageQuestionsToProto(request *http.Request) *session.GetQuestionsByPageRequest {
	return &session.GetQuestionsByPageRequest{
		Page: request.URL.Query()["p"][0],
	}
}

func convertAddStatisticsToProto(request *http.Request) (*session.AddStatisticsRequest, error) {
	var addStatistics *addStatisticsRequest
	err := json.NewDecoder(request.Body).Decode(&addStatistics)
	if err != nil {
		return nil, err
	}

	protoStatistics := &session.AddStatisticsRequest{
		Page: addStatistics.Page,
	}

	newStatistics := make([]*session.NewStatistics, 0, len(addStatistics.Statistics))
	for _, stat := range addStatistics.Statistics {
		newStatistics = append(newStatistics, &session.NewStatistics{
			Uuid:         stat.Uuid,
			IsAdditional: stat.IsAdditional,
			Score:        stat.Score,
		})
	}

	protoStatistics.NewStatistics = newStatistics

	return protoStatistics, nil
}

func convertGetStatisticsByPageToProto(request *http.Request) *session.GetStatisticsRequest {
	return &session.GetStatisticsRequest{
		Page: request.URL.Query()["p"][0],
	}
}
