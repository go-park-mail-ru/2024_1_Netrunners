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
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/users/service"
)

type UserPageHandlers struct {
	usersService   service.UsersService
	sessionService SessionService
	logger         *zap.SugaredLogger
}

func NewUserPageHandlers(usersService service.UsersService, sessionService SessionService,
	logger *zap.SugaredLogger) *UserPageHandlers {
	return &UserPageHandlers{
		usersService:   usersService,
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
	user, err := UserPageHandlers.usersService.GetUserDataByUuid(ctx, uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
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
	userPreview, err := UserPageHandlers.usersService.GetUserPreview(ctx, uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	user, err := UserPageHandlers.usersService.GetUserDataByUuid(ctx, uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

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

	_, err = UserPageHandlers.usersService.IsTokenValid(userToken)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	uuid := mux.Vars(r)["uuid"]
	currUser, err := UserPageHandlers.usersService.GetUserDataByUuid(ctx, uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

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

		currUser, err = UserPageHandlers.usersService.ChangeUserPasswordByUuid(ctx, uuid, newData)
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
			}
			return
		}

	case "chUsername":
		err = service.ValidateUsername(newData)
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
			}
			return
		}

		currUser, err = UserPageHandlers.usersService.ChangeUserNameByUuid(ctx, uuid, newData)
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
			}
			return
		}
	case "chAvatar":
		files := r.MultipartForm.File["avatar"]
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
			}
			return
		}

		err = UserPageHandlers.saveFile(files[0], currUser.Email, requestId.(string))
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
			}
			return
		}
	}

	version := currUser.Version + 1

	tokenSigned, err := UserPageHandlers.usersService.GenerateTokens(currUser.Email, false, version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	err = UserPageHandlers.sessionService.Add(ctx, currUser.Email, tokenSigned, version)
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
