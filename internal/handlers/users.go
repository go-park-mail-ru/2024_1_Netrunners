package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
)

type UserPageHandlers struct {
	authService *service.AuthService
	logger      *zap.SugaredLogger
}

func NewUserPageHandlers(authService *service.AuthService, logger *zap.SugaredLogger) *UserPageHandlers {
	return &UserPageHandlers{
		authService: authService,
		logger:      logger,
	}
}

type profileResponse struct {
	Status   int         `json:"status"`
	UserInfo domain.User `json:"user"`
}

func (UserPageHandlers *UserPageHandlers) GetProfileData(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]
	requestID := generateRequestID()
	ctx := reqIdCTX(requestID)

	user, err := UserPageHandlers.authService.GetUserDataByUuid(ctx, uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] error at writing response: %v\n", requestID, err)
		}
		return
	}

	response := profileResponse{
		Status:   http.StatusOK,
		UserInfo: user,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] error at writing response: %v\n", requestID, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] error at writing response: %v\n", requestID, err)
		}
		return
	}
}

type profilePreviewResponse struct {
	Status      int                `json:"status"`
	UserPreview domain.UserPreview `json:"user"`
}

func (UserPageHandlers *UserPageHandlers) GetProfilePreview(w http.ResponseWriter, r *http.Request) {
	requestID := generateRequestID()
	ctx := reqIdCTX(requestID)

	var inputUserData domain.User
	err := json.NewDecoder(r.Body).Decode(&inputUserData)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] error at writing response: %v\n", requestID, err)
		}
		return
	}

	uuid := inputUserData.Uuid
	userPreview, err := UserPageHandlers.authService.GetUserPreview(ctx, uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] error at writing response: %v\n", requestID, err)
		}
		return
	}

	response := profilePreviewResponse{
		Status:      http.StatusOK,
		UserPreview: userPreview,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] error at writing response: %v\n", requestID, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] error at writing response: %v\n", requestID, err)
		}
		return
	}
}
