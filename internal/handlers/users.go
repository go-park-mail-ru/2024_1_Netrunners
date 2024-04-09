package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
)

type UserPageHandlers struct {
	authService    *service.AuthService
	sessionService *service.SessionService
	logger         *zap.SugaredLogger
}

func NewUserPageHandlers(authService *service.AuthService, sessionService *service.SessionService,
	logger *zap.SugaredLogger) *UserPageHandlers {
	return &UserPageHandlers{
		authService:    authService,
		sessionService: sessionService,
		logger:         logger,
	}
}

type profileResponse struct {
	Status   int         `json:"status"`
	UserInfo domain.User `json:"user"`
}

func (UserPageHandlers *UserPageHandlers) GetProfileData(w http.ResponseWriter, r *http.Request) {
	requestID := requestId.GenerateRequestID()
	ctx := requestId.GenerateReqIdCTX(requestID)

	uuid := mux.Vars(r)["uuid"]
	user, err := UserPageHandlers.authService.GetUserDataByUuid(ctx, uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	escapeUserData(&user)
	response := profileResponse{
		Status:   http.StatusOK,
		UserInfo: user,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}

type profilePreviewResponse struct {
	Status      int                `json:"status"`
	UserPreview domain.UserPreview `json:"user"`
}

func (UserPageHandlers *UserPageHandlers) GetProfilePreview(w http.ResponseWriter, r *http.Request) {
	requestID := requestId.GenerateRequestID()
	ctx := requestId.GenerateReqIdCTX(requestID)

	var inputUserData domain.User
	err := json.NewDecoder(r.Body).Decode(&inputUserData)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	uuid := inputUserData.Uuid
	userPreview, err := UserPageHandlers.authService.GetUserPreview(ctx, uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	escapeUserPreviewData(&userPreview)

	response := profilePreviewResponse{
		Status:      http.StatusOK,
		UserPreview: userPreview,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
}

func (UserPageHandlers *UserPageHandlers) ProfileEditByUuid(w http.ResponseWriter, r *http.Request) {
	requestID := requestId.GenerateRequestID()
	ctx := requestId.GenerateReqIdCTX(requestID)

	userToken, err := r.Cookie("access")
	if err != nil {
		err = WriteError(w, myerrors.ErrNoActiveSession)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	_, err = UserPageHandlers.authService.IsTokenValid(userToken)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	uuid := mux.Vars(r)["uuid"]

	currUser, err := UserPageHandlers.authService.GetUserDataByUuid(ctx, uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	action := r.FormValue("action")
	if action == "chPassword" {
		newPassword := r.FormValue("newData")
		err = service.ValidatePassword(newPassword)
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
			}
			return
		}

		currUser, err = UserPageHandlers.authService.ChangeUserPasswordByUuid(ctx, uuid, newPassword)
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
			}
			return
		}
	}

	if action == "chUsername" {
		newUsername := r.FormValue("newData")
		err = service.ValidateUsername(newUsername)
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
			}
			return
		}

		currUser, err = UserPageHandlers.authService.ChangeUserNameByUuid(ctx, uuid, newUsername)
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
			}
			return
		}
	}

	if action == "chAvatar" {
		files := r.MultipartForm.File["avatar"]
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
			}
			return
		}

		for _, file := range files {
			src, err := file.Open()
			if err != nil {
				err = WriteError(w, err)
				if err != nil {
					UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
				}
				return
			}
			defer src.Close()

			dst, err := os.Create("./avatars/" + file.Filename)
			if err != nil {
				err = WriteError(w, err)
				if err != nil {
					UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
				}
				return
			}
			defer dst.Close()

			if _, err := io.Copy(dst, src); err != nil {
				err = WriteError(w, err)
				if err != nil {
					UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
				}
				return
			}

			UserPageHandlers.logger.Errorf("[reqid=%s] avatar uploaded successfully\n", requestID)
		}
	}

	version := currUser.Version + 1

	tokenSigned, err := UserPageHandlers.authService.GenerateTokens(currUser.Email, false, version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	err = UserPageHandlers.sessionService.Add(ctx, currUser.Email, tokenSigned, version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	tokenCookie := &http.Cookie{
		Name:     "access",
		Value:    tokenSigned,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   tokenCookieExpirationTime,
	}

	http.SetCookie(w, tokenCookie)

	err = WriteSuccess(w)
	if err != nil {
		UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
	}

}
