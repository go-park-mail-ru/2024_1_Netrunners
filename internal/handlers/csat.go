package handlers

import (
	"net/http"

	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
	"go.uber.org/zap"
)

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

}

func (csatHandlers *CsatHandlers) AddStatistics(w http.ResponseWriter, r *http.Request) {

}

func (csatHandlers *CsatHandlers) GetStatisticsByPage(w http.ResponseWriter, r *http.Request) {

}
