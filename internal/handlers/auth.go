package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/metrics"
	reqid "github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
)

var (
	tokenCookieExpirationTime = 48 * 3600
)

type AuthService interface {
	CreateUser(ctx context.Context, user domain.UserSignUp) error
	RemoveUser(ctx context.Context, email string) error
	HasUser(ctx context.Context, email, password string) error
	GetUser(ctx context.Context, email string) (domain.User, error)
	ChangeUserPassword(ctx context.Context, email, newPassword string) (domain.User, error)
	ChangeUserName(ctx context.Context, email, newName string) (domain.User, error)
	GetUserDataByUuid(ctx context.Context, uuid string) (domain.User, error)
	GetUserPreview(ctx context.Context, uuid string) (domain.UserPreview, error)
	ChangeUserPasswordByUuid(ctx context.Context, uuid, newPassword string) (domain.User, error)
	ChangeUserNameByUuid(ctx context.Context, uuid, newName string) (domain.User, error)
	IsTokenValid(token *http.Cookie) (jwt.MapClaims, error)
	GenerateTokens(login string, isAdmin bool, version uint32) (tokenSigned string, err error)
}

type SessionService interface {
	Add(ctx context.Context, login string, token string, version uint32) (err error)
	DeleteSession(ctx context.Context, login string, token string) (err error)
	Update(ctx context.Context, login string, token string) (err error)
	CheckVersion(ctx context.Context, login string, token string, usersVersion uint32) (hasSession bool, err error)
	GetVersion(ctx context.Context, login string, token string) (version uint32, err error)
	HasSession(ctx context.Context, login string, token string) error
	CheckAllUserSessionTokens(ctx context.Context, login string) error
	GenerateTokens(login string, isAdmin bool, version uint32) (tokenSigned string, err error)
	IsTokenValid(token *http.Cookie) (jwt.MapClaims, error)
}

type AuthPageHandlers struct {
	usersClient    *session.UsersClient
	sessionsClient *session.SessionsClient
	metrics        *metrics.HttpMetrics
	logger         *zap.SugaredLogger
}

func NewAuthPageHandlers(usersClient *session.UsersClient, sessionsClient *session.SessionsClient,
	metrics *metrics.HttpMetrics, logger *zap.SugaredLogger) *AuthPageHandlers {
	return &AuthPageHandlers{
		usersClient:    usersClient,
		sessionsClient: sessionsClient,
		metrics:        metrics,
		logger:         logger,
	}
}

func (authPageHandlers *AuthPageHandlers) Login(w http.ResponseWriter, r *http.Request) {
	var inputUserData domain.UserSignUp
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	err := easyjson.UnmarshalFromReader(r.Body, &inputUserData)
	if err != nil {
		authPageHandlers.logger.Errorf("[reqid=%s] failed to decode: %v\n", requestID, myerrors.ErrFailedDecode)
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
	login := inputUserData.Email
	password := inputUserData.Password

	err = ValidateLogin(login)
	if err != nil {
		authPageHandlers.logger.Errorf("[reqid=%s] login is not valid: %v\n", requestID,
			myerrors.ErrLoginIsNotValid)
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	err = ValidatePassword(password)
	if err != nil {
		authPageHandlers.logger.Errorf("[reqid=%s] password is not valid: %v\n", requestID,
			myerrors.ErrPasswordIsToShort)
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	req := session.HasUserRequest{Login: login, Password: password}
	has, err := (*authPageHandlers.usersClient).HasUser(ctx, &req)
	if err != nil {
		authPageHandlers.logger.Errorf("[reqid=%s] failed to check sessions: %v\n", requestID,
			myerrors.ErrNoActiveSession)
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
	if has.Has {
		authPageHandlers.logger.Errorf("[reqid=%s] user already exists: %v\n", requestID,
			myerrors.ErrUserAlreadyExists)
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	reqGetUser := session.GetUserRequest{Login: login}
	user, err := (*authPageHandlers.usersClient).GetUser(ctx, &reqGetUser)
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	tokenSigned, err := GenerateTokens(user.User.Email, user.User.IsAdmin, user.User.Version)
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	reqAdd := session.AddRequest{Login: user.User.Email, Token: tokenSigned, Version: user.User.Version}
	_, err = (*authPageHandlers.sessionsClient).Add(ctx, &reqAdd)
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}

		return
	}

	uuidCookie := &http.Cookie{
		Name:     "user_uuid",
		Value:    user.User.Uuid,
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
		MaxAge:   0,
	}
	http.SetCookie(w, uuidCookie)

	tokenCookie := &http.Cookie{
		Name:     "access",
		Value:    tokenSigned,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   tokenCookieExpirationTime,
	}

	http.SetCookie(w, tokenCookie)
	err = WriteSuccess(w, r, authPageHandlers.metrics)
	if err != nil {
		authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
	}

	authPageHandlers.logger.Info(fmt.Sprintf("[reqid=%s] success login", requestID))
}

func (authPageHandlers *AuthPageHandlers) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	userToken, err := r.Cookie("access")
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	tokenClaims, err := IsTokenValid(userToken, os.Getenv("SECRETKEY"))
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	reqDel := session.DeleteSessionRequest{Login: tokenClaims["Login"].(string), Token: userToken.Value}
	_, err = (*authPageHandlers.sessionsClient).DeleteSession(ctx, &reqDel)
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
	tokenCookie := &http.Cookie{
		Name:     "access",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   0,
	}

	http.SetCookie(w, tokenCookie)

	reqCheck := session.GetVersionRequest{Login: tokenClaims["Login"].(string), Token: userToken.Value}
	_, err = (*authPageHandlers.sessionsClient).GetVersion(ctx, &reqCheck)
	if err != nil {
		authPageHandlers.logger.Info(fmt.Sprintf("[reqid=%s] success logout", requestID))
	}

	err = WriteSuccess(w, r, authPageHandlers.metrics)
	if err != nil {
		authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
	}
}

func (authPageHandlers *AuthPageHandlers) Signup(w http.ResponseWriter, r *http.Request) {
	var inputUserData domain.UserSignUp
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	err := easyjson.UnmarshalFromReader(r.Body, &inputUserData)
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	login := inputUserData.Email
	username := inputUserData.Name
	password := inputUserData.Password

	err = ValidateLogin(login)
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	err = ValidateUsername(username)
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	err = ValidatePassword(password)
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	var version uint32 = 1

	var user = domain.UserSignUp{
		Email:    login,
		Name:     username,
		Password: password,
	}

	reqCreate := session.CreateUserRequest{User: convertUserSignUpDataToRegular(user)}
	_, err = (*authPageHandlers.usersClient).CreateUser(ctx, &reqCreate)
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	tokenSigned, err := GenerateTokens(user.Email, false, version)
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	reqAdd := session.AddRequest{Login: user.Email, Token: tokenSigned}
	_, err = (*authPageHandlers.sessionsClient).Add(ctx, &reqAdd)
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	reqGetUser := session.GetUserRequest{Login: user.Email}
	userForUuid, err := (*authPageHandlers.usersClient).GetUser(ctx, &reqGetUser)
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	uuidCookie := &http.Cookie{
		Name:     "user_uuid",
		Value:    userForUuid.User.Uuid,
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
		MaxAge:   0,
	}
	http.SetCookie(w, uuidCookie)

	tokenCookie := &http.Cookie{
		Name:     "access",
		Value:    tokenSigned,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   tokenCookieExpirationTime,
	}

	http.SetCookie(w, tokenCookie)

	err = WriteSuccess(w, r, authPageHandlers.metrics)
	if err != nil {
		authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
	}
}

func (authPageHandlers *AuthPageHandlers) Check(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	userToken, err := r.Cookie("access")
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, myerrors.ErrNoActiveSession)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	tokenClaims, err := IsTokenValid(userToken, os.Getenv("SECRETKEY"))
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	reqHas := session.HasSessionRequest{Login: tokenClaims["Login"].(string), Token: userToken.Value}
	_, err = (*authPageHandlers.sessionsClient).HasSession(ctx, &reqHas)
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, myerrors.ErrNoActiveSession)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	tokenSigned, err := GenerateTokens(tokenClaims["Login"].(string), tokenClaims["IsAdmin"].(bool),
		uint32(tokenClaims["Version"].(float64)))
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	reqAdd := session.AddRequest{Login: tokenClaims["Login"].(string), Token: tokenSigned,
		Version: uint32(tokenClaims["Version"].(float64))}
	_, err = (*authPageHandlers.sessionsClient).Add(ctx, &reqAdd)
	if err != nil {
		err = WriteError(w, r, authPageHandlers.metrics, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	tokenCookie := &http.Cookie{
		Name:     "access",
		Value:    tokenSigned,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   tokenCookieExpirationTime,
	}

	http.SetCookie(w, tokenCookie)
	err = WriteSuccess(w, r, authPageHandlers.metrics)
	if err != nil {
		authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
	}
}
