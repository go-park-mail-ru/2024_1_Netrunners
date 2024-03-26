package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
)

var (
	accessCookieExpirationTime  = 5 * 60
	refreshCookieExpirationTime = 48 * 3600
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
	var inputUserData domain.User
	err := json.NewDecoder(r.Body).Decode(&inputUserData)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			if err != nil {
				authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
			}
		}
		return
	}
	login := inputUserData.Email
	password := inputUserData.Password

	err = service.ValidateLogin(login)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	err = service.ValidatePassword(password)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	err = authPageHandlers.authService.HasUser(login, password)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	user, err := authPageHandlers.authService.GetUser(login)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	accessTokenSigned, refreshTokenSigned, err := authPageHandlers.authService.GenerateTokens(login,
		user.IsAdmin, user.Version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	err = authPageHandlers.sessionService.Add(user.Email, refreshTokenSigned, user.Version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}

		return
	}

	accessCookie := &http.Cookie{
		Name:     "access",
		Value:    accessTokenSigned,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   accessCookieExpirationTime,
	}

	refreshCookie := &http.Cookie{
		Name:     "refresh",
		Value:    refreshTokenSigned,
		Path:     "/auth",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   refreshCookieExpirationTime,
	}

	http.SetCookie(w, refreshCookie)
	http.SetCookie(w, accessCookie)
	err = WriteSuccess(w)
	if err != nil {
		authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
	}

	authPageHandlers.logger.Info("success login")
}

func (authPageHandlers *AuthPageHandlers) Logout(w http.ResponseWriter, r *http.Request) {
	userRefreshToken, err := r.Cookie("refresh")
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	refreshTokenClaims, err := authPageHandlers.authService.IsTokenValid(userRefreshToken)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	err = authPageHandlers.sessionService.DeleteSession(refreshTokenClaims["Login"].(string), userRefreshToken.Value)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	accessCookie := &http.Cookie{
		Name:     "access",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   0,
	}

	refreshCookie := &http.Cookie{
		Name:     "refresh",
		Value:    "",
		Path:     "/auth",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   0,
	}

	http.SetCookie(w, refreshCookie)
	http.SetCookie(w, accessCookie)

	_, err = authPageHandlers.sessionService.GetVersion(refreshTokenClaims["Login"].(string), userRefreshToken.Value)
	if err != nil {
		if err != nil {
			authPageHandlers.logger.Info("success logout")
		}
	}

	err = WriteSuccess(w)
	if err != nil {
		authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
	}
}

func (authPageHandlers *AuthPageHandlers) Signup(w http.ResponseWriter, r *http.Request) {
	var inputUserData domain.UserSignUp

	err := json.NewDecoder(r.Body).Decode(&inputUserData)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
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
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	err = service.ValidateUsername(username)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	err = service.ValidatePassword(password)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	var version uint8 = 1

	var user = domain.UserSignUp{
		Email:    login,
		Name:     username,
		Password: password,
	}

	err = authPageHandlers.authService.CreateUser(user)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	accessTokenSigned, refreshTokenSigned, err :=
		authPageHandlers.authService.GenerateTokens(login, false, version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	err = authPageHandlers.sessionService.Add(login, refreshTokenSigned, version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}

		return
	}

	accessCookie := &http.Cookie{
		Name:     "access",
		Value:    accessTokenSigned,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   accessCookieExpirationTime,
	}

	refreshCookie := &http.Cookie{
		Name:     "refresh",
		Value:    refreshTokenSigned,
		Path:     "/auth",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   refreshCookieExpirationTime,
	}

	http.SetCookie(w, refreshCookie)
	http.SetCookie(w, accessCookie)

	err = WriteSuccess(w)
	if err != nil {
		authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
	}
}

func (authPageHandlers *AuthPageHandlers) Check(w http.ResponseWriter, r *http.Request) {
	userRefreshToken, err := r.Cookie("refresh")
	if err != nil {
		err = WriteError(w, myerrors.ErrNoActiveSession)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	refreshTokenClaims, err := authPageHandlers.authService.IsTokenValid(userRefreshToken)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	err = authPageHandlers.sessionService.HasSession(refreshTokenClaims["Login"].(string), userRefreshToken.Value)
	if err != nil {
		err = WriteError(w, myerrors.ErrNoActiveSession)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	userAccessToken, err := r.Cookie("access")
	if err == nil {
		_, err = authPageHandlers.authService.IsTokenValid(userAccessToken)
		if err == nil {
			err = WriteSuccess(w)
			if err != nil {
				authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
			}
			return
		}
	}

	accessTokenSigned, refreshTokenSigned, err := authPageHandlers.authService.GenerateTokens(
		refreshTokenClaims["Login"].(string),
		refreshTokenClaims["IsAdmin"].(bool),
		uint8(refreshTokenClaims["Version"].(float64)))
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	err = authPageHandlers.sessionService.Add(refreshTokenClaims["Login"].(string), refreshTokenSigned,
		uint8(refreshTokenClaims["Version"].(float64)))
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	accessCookie := &http.Cookie{
		Name:     "access",
		Value:    accessTokenSigned,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   accessCookieExpirationTime,
	}

	refreshCookie := &http.Cookie{
		Name:     "refresh",
		Value:    refreshTokenSigned,
		Path:     "/auth",
		HttpOnly: true,
		MaxAge:   refreshCookieExpirationTime,
	}

	http.SetCookie(w, refreshCookie)
	http.SetCookie(w, accessCookie)
	err = WriteSuccess(w)
	if err != nil {
		authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
	}
}
