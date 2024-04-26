package handlers

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
	reqid "github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
)

type UserPageHandlers struct {
	client         *session.UsersClient
	sessionService SessionService
	logger         *zap.SugaredLogger
}

func NewUserPageHandlers(client *session.UsersClient, sessionService SessionService,
	logger *zap.SugaredLogger) *UserPageHandlers {
	return &UserPageHandlers{
		client:         client,
		sessionService: sessionService,
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
	userProto, err := (*UserPageHandlers.client).GetUserDataByUuid(ctx, &req)
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
	userPreviewProto, err := (*UserPageHandlers.client).GetUserPreview(ctx, &req)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}
	userPreview := convertUserPreviewToRegular(userPreviewProto.User)

	reqData := session.GetUserDataByUuidRequest{Uuid: uuid}
	userProto, err := (*UserPageHandlers.client).GetUserDataByUuid(ctx, &reqData)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}
	user := convertUserToRegular(userProto.User)

	avatar := "./uploads/users/" + user.Email + "/avatar.png"
	_, err = os.Stat(avatar)
	if err != nil {
		userPreview.Avatar = "./uploads/users/default/avatar.png"
	} else {
		userPreview.Avatar = avatar
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

	_, err = UserPageHandlers.sessionService.IsTokenValid(userToken)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	uuid := mux.Vars(r)["uuid"]
	req := session.GetUserDataByUuidRequest{Uuid: uuid}
	getUserByDataRes, err := (*UserPageHandlers.client).GetUserDataByUuid(ctx, &req)
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
		err = service.ValidatePassword(newData)
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
			}
			return
		}

		reqPass := session.ChangeUserPasswordByUuidRequest{Uuid: uuid, NewPassword: newData}
		changePassRes, err := (*UserPageHandlers.client).ChangeUserPasswordByUuid(ctx, &reqPass)
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
			}
			return
		}
		currUserProto = changePassRes.User

	case "chUsername":
		err = service.ValidateUsername(newData)
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
			}
			return
		}

		reqName := session.ChangeUserNameByUuidRequest{Uuid: uuid, NewUsername: newData}
		changeNameRes, err := (*UserPageHandlers.client).ChangeUserNameByUuid(ctx, &reqName)
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
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
			}
			return
		}

		err = UserPageHandlers.saveFile(files[0], currUserProto.Email, requestId.(string))
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
			}
			return
		}
	}

	version := currUserProto.Version + 1

	tokenSigned, err := UserPageHandlers.sessionService.GenerateTokens(currUserProto.Email, false, version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	err = UserPageHandlers.sessionService.Add(ctx, currUserProto.Email, tokenSigned, version)
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

func (UserPageHandlers *UserPageHandlers) saveFile(file *multipart.FileHeader, email string, requestId string) error {
	storagePath := "./uploads/users/" + email
	_, err := os.Stat(storagePath)
	if err != nil {
		err = os.Mkdir(storagePath, 0755)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to create directory for user's avatar: %v\n", requestId, err)
			return myerrors.ErrInternalServerError
		}
	}

	dst, err := os.Create(storagePath + "/avatar.png")
	if err != nil {
		UserPageHandlers.logger.Errorf("[reqid=%s] failed to create empty avatar file: %v\n", requestId, err)
		return myerrors.ErrInternalServerError
	}

	defer dst.Close()

	src, err := file.Open()
	if err != nil {
		UserPageHandlers.logger.Errorf("[reqid=%s] failed to open created avatar file: %v\n", requestId, err)
		return myerrors.ErrInternalServerError
	}
	defer src.Close()

	if _, err := io.Copy(dst, src); err != nil {
		UserPageHandlers.logger.Errorf("[reqid=%s] failed to copy into avatar file: %v\n", requestId, err)
		return myerrors.ErrInternalServerError
	}

	return nil
}
