package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
	reqid "github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
)

var (
	tokenCookieExpirationTime = 48 * 3600
)

type AuthPageHandlers struct {
	authService    *service.AuthService
	sessionService *service.SessionService
	logger         *zap.SugaredLogger
}

func NewAuthPageHandlers(authService *service.AuthService, sessionService *service.SessionService,
	logger *zap.SugaredLogger) *AuthPageHandlers {
	return &AuthPageHandlers{
		authService:    authService,
		sessionService: sessionService,
		logger:         logger,
	}
}

func (authPageHandlers *AuthPageHandlers) Login(w http.ResponseWriter, r *http.Request) {
	var inputUserData domain.UserSignUp
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	err := json.NewDecoder(r.Body).Decode(&inputUserData)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}
	login := inputUserData.Email
	password := inputUserData.Password

	err = service.ValidateLogin(login)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	err = service.ValidatePassword(password)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	err = authPageHandlers.authService.HasUser(ctx, login, password)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	user, err := authPageHandlers.authService.GetUser(ctx, login)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	tokenSigned, err := authPageHandlers.authService.GenerateTokens(login, user.IsAdmin, user.Version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	err = authPageHandlers.sessionService.Add(ctx, user.Email, tokenSigned, user.Version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}

		return
	}

	uuidCookie := &http.Cookie{
		Name:     "user_uuid",
		Value:    user.Uuid,
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

	csrfToken, err := generateCsrfToken()
	if err != nil {
		authPageHandlers.logger.Errorf("[reqid=%s] failed to generate csrf: %v\n", requestID, err)
		return
	}
	w.Header().Set("X-CSRF-TOKEN", csrfToken)
	err = WriteSuccess(w)
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
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	tokenClaims, err := authPageHandlers.authService.IsTokenValid(userToken)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	err = authPageHandlers.sessionService.DeleteSession(ctx, tokenClaims["Login"].(string), userToken.Value)
	if err != nil {
		err = WriteError(w, err)
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

	_, err = authPageHandlers.sessionService.GetVersion(ctx, tokenClaims["Login"].(string), userToken.Value)
	if err != nil {
		authPageHandlers.logger.Info(fmt.Sprintf("[reqid=%s] success logout", requestID))
	}

	err = WriteSuccess(w)
	if err != nil {
		authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
	}
}

func (authPageHandlers *AuthPageHandlers) Signup(w http.ResponseWriter, r *http.Request) {
	var inputUserData domain.UserSignUp
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	err := json.NewDecoder(r.Body).Decode(&inputUserData)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	login := inputUserData.Email
	username := inputUserData.Name
	password := inputUserData.Password

	err = service.ValidateLogin(login)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	err = service.ValidateUsername(username)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	err = service.ValidatePassword(password)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	var version uint8 = 1

	var user = domain.UserSignUp{
		Email:    login,
		Name:     username,
		Password: password,
	}

	err = authPageHandlers.authService.CreateUser(ctx, user)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	tokenSigned, err := authPageHandlers.authService.GenerateTokens(login, false, version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	err = authPageHandlers.sessionService.Add(ctx, login, tokenSigned, version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	userForUuid, err := authPageHandlers.authService.GetUser(ctx, login)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	uuidCookie := &http.Cookie{
		Name:     "user_uuid",
		Value:    userForUuid.Uuid,
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

	csrfToken, err := generateCsrfToken()
	if err != nil {
		authPageHandlers.logger.Errorf("[reqid=%s] failed to generate csrf: %v\n", requestID, err)
		return
	}
	w.Header().Set("X-CSRF-TOKEN", csrfToken)
	err = WriteSuccess(w)
	if err != nil {
		authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
	}
}

func (authPageHandlers *AuthPageHandlers) Check(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value(reqid.ReqIDKey)

	userToken, err := r.Cookie("access")
	if err != nil {
		err = WriteError(w, myerrors.ErrNoActiveSession)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	tokenClaims, err := authPageHandlers.authService.IsTokenValid(userToken)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	err = authPageHandlers.sessionService.HasSession(ctx, tokenClaims["Login"].(string), userToken.Value)
	if err != nil {
		err = WriteError(w, myerrors.ErrNoActiveSession)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	tokenSigned, err := authPageHandlers.authService.GenerateTokens(
		tokenClaims["Login"].(string),
		tokenClaims["IsAdmin"].(bool),
		uint8(tokenClaims["Version"].(float64)))
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
		}
		return
	}

	err = authPageHandlers.sessionService.Add(ctx, tokenClaims["Login"].(string), tokenSigned,
		uint8(tokenClaims["Version"].(float64)))
	if err != nil {
		err = WriteError(w, err)
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

	csrfToken, err := generateCsrfToken()
	if err != nil {
		authPageHandlers.logger.Errorf("[reqid=%s] failed to generate csrf: %v\n", requestID, err)
		return
	}
	w.Header().Set("X-CSRF-TOKEN", csrfToken)
	err = WriteSuccess(w)
	if err != nil {
		authPageHandlers.logger.Errorf("[reqid=%s] failed to write response: %v\n", requestID, err)
	}
}

func generateCsrfToken() (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		Issuer:    "NETrunnerFLIX",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenSigned, err := token.SignedString([]byte("sad"))
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	return tokenSigned, nil
}
