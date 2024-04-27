package handlers

import (
	"encoding/json"
	"net/http"

	reqid "github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
	"go.uber.org/zap"
)

type getPageQuestionsResponse struct {
}

type addStatisticsResponse struct {
}

type getStatisticsByPageResponse struct {
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

	_, err := (*csatHandlers.csatClient).GetQuestionsByPage(ctx, convertGetPageQuestionsToProto(r))
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			csatHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	response := getPageQuestionsResponse{}

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

	_, err := (*csatHandlers.csatClient).AddStatistics(ctx, convertAddStatisticsToProto(r))
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			csatHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	response := getPageQuestionsResponse{}

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

	_, err := (*csatHandlers.csatClient).GetStatisticsByPage(ctx, convertGetStatisticsByPageToProto(r))
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			csatHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	response := getPageQuestionsResponse{}

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

func convertAddStatisticsToProto(request *http.Request) *session.AddStatisticsRequest {
	return &session.AddStatisticsRequest{}
}

func convertGetStatisticsByPageToProto(request *http.Request) *session.GetStatisticsRequest {
	return &session.GetStatisticsRequest{}
}
