package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
	reqid "github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
)

type UserPageHandlers struct {
	usersClient    *session.UsersClient
	sessionsClient *session.SessionsClient
	logger         *zap.SugaredLogger
}

func NewUserPageHandlers(usersClient *session.UsersClient, sessionsClient *session.SessionsClient,
	logger *zap.SugaredLogger) *UserPageHandlers {
	return &UserPageHandlers{
		usersClient:    usersClient,
		sessionsClient: sessionsClient,
		logger:         logger,
	}
}

type profileResponse struct {
	Status   int         `json:"status"`
	UserInfo domain.User `json:"user"`
}

func (UserPageHandlers *UserPageHandlers) GetProfileData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := ctx.Value(reqid.ReqIDKey)
	uuid := mux.Vars(r)["uuid"]
	req := session.GetUserDataByUuidRequest{Uuid: uuid}
	userProto, err := (*UserPageHandlers.usersClient).GetUserDataByUuid(ctx, &req)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}
	user := convertUserToRegular(userProto.User)

	escapeUserData(&user)
	response := profileResponse{
		Status:   http.StatusOK,
		UserInfo: user,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}
}

type profilePreviewResponse struct {
	Status      int                `json:"status"`
	UserPreview domain.UserPreview `json:"user"`
}

func (UserPageHandlers *UserPageHandlers) GetProfilePreview(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := ctx.Value(reqid.ReqIDKey)

	uuid := mux.Vars(r)["uuid"]
	req := session.GetUserPreviewRequest{Uuid: uuid}
	userPreviewProto, err := (*UserPageHandlers.usersClient).GetUserPreview(ctx, &req)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}
	userPreview := convertUserPreviewToRegular(userPreviewProto.User)

	escapeUserPreviewData(&userPreview)

	response := profilePreviewResponse{
		Status:      http.StatusOK,
		UserPreview: userPreview,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}
}

func (UserPageHandlers *UserPageHandlers) ProfileEditByUuid(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := ctx.Value(reqid.ReqIDKey)

	userToken, err := r.Cookie("access")
	if err != nil {
		err = WriteError(w, myerrors.ErrNoActiveSession)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	_, err = IsTokenValid(userToken, os.Getenv("SECRETKEY"))
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	uuid := mux.Vars(r)["uuid"]
	req := session.GetUserDataByUuidRequest{Uuid: uuid}
	getUserByDataRes, err := (*UserPageHandlers.usersClient).GetUserDataByUuid(ctx, &req)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}
	currUserProto := getUserByDataRes.User

	newData := r.FormValue("newData")
	switch r.FormValue("action") {
	case "chPassword":
		err = ValidatePassword(newData)
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
			}
			return
		}

		reqPass := session.ChangeUserPasswordByUuidRequest{Uuid: uuid, NewPassword: newData}
		changePassRes, err := (*UserPageHandlers.usersClient).ChangeUserPasswordByUuid(ctx, &reqPass)
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
			}
			return
		}
		currUserProto = changePassRes.User

	case "chUsername":
		err = ValidateUsername(newData)
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
			}
			return
		}

		reqName := session.ChangeUserNameByUuidRequest{Uuid: uuid, NewUsername: newData}
		changeNameRes, err := (*UserPageHandlers.usersClient).ChangeUserNameByUuid(ctx, &reqName)
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
			}
			return
		}
		currUserProto = changeNameRes.User
	case "chAvatar":
		files := r.MultipartForm.File["avatar"]

		avatar64, err := Encode(files[0])
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to encode into base64: %v\n", requestId, err)
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
			}
			return
		}

		reqAvatar := session.ChangeUserAvatarByUuidRequest{Uuid: uuid, NewAvatar: avatar64}
		changeAvatarRes, err := (*UserPageHandlers.usersClient).ChangeUserAvatarByUuid(ctx, &reqAvatar)
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
			}
			return
		}
		currUserProto = changeAvatarRes.User
	}

	version := currUserProto.Version + 1

	tokenSigned, err := GenerateTokens(currUserProto.Email, false, version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	reqAdd := session.AddRequest{Login: currUserProto.Email, Token: tokenSigned, Version: version}
	_, err = (*UserPageHandlers.sessionsClient).Add(ctx, &reqAdd)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
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
		UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
	}
}
