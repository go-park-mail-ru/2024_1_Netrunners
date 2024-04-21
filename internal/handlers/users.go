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
	httpctx "github.com/go-park-mail-ru/2024_1_Netrunners/internal/httpcontext"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
)

type UserPageHandlers struct {
	authService    AuthService
	sessionService SessionService
	logger         *zap.SugaredLogger
}

func NewUserPageHandlers(authService AuthService, sessionService SessionService,
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
	ctx := r.Context()
	requestId := ctx.Value(httpctx.ReqIDKey)
	ctxEmail := ctx.Value(httpctx.CtxEmail).(string)
	uuid := mux.Vars(r)["uuid"]

	currentUser, err := UserPageHandlers.authService.GetUser(ctx, ctxEmail)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	if currentUser.Uuid != uuid {
		err = WriteError(w, myerrors.ErrForbidden)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	escapeUserData(&currentUser)
	response := profileResponse{
		Status:   http.StatusOK,
		UserInfo: currentUser,
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
	requestId := ctx.Value(httpctx.ReqIDKey)
	ctxEmail := ctx.Value(httpctx.CtxEmail).(string)
	uuid := mux.Vars(r)["uuid"]

	currentUser, err := UserPageHandlers.authService.GetUser(ctx, ctxEmail)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	if currentUser.Uuid != uuid {
		err = WriteError(w, myerrors.ErrForbidden)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	userPreview, err := UserPageHandlers.authService.GetUserPreview(ctx, uuid)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	avatar := "./uploads/users/" + currentUser.Email + "/avatar.png"
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
	requestId := ctx.Value(httpctx.ReqIDKey)

	userToken, err := r.Cookie("access")
	if err != nil {
		err = WriteError(w, myerrors.ErrNoActiveSession)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	_, err = UserPageHandlers.authService.IsTokenValid(userToken)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	ctxEmail := ctx.Value(httpctx.CtxEmail).(string)
	uuid := mux.Vars(r)["uuid"]

	currentUser, err := UserPageHandlers.authService.GetUser(ctx, ctxEmail)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	if currentUser.Uuid != uuid {
		err = WriteError(w, myerrors.ErrForbidden)
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

		currentUser, err = UserPageHandlers.authService.ChangeUserPasswordByUuid(ctx, uuid, newData)
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

		currentUser, err = UserPageHandlers.authService.ChangeUserNameByUuid(ctx, uuid, newData)
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

		err = UserPageHandlers.saveFile(files[0], currentUser.Email, requestId.(string))
		if err != nil {
			err = WriteError(w, err)
			if err != nil {
				UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
			}
			return
		}
	}

	version := currentUser.Version + 1

	tokenSigned, err := UserPageHandlers.authService.GenerateTokens(currentUser.Email, false, version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			UserPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestId, err)
		}
		return
	}

	err = UserPageHandlers.sessionService.Add(ctx, currentUser.Email, tokenSigned, version)
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
