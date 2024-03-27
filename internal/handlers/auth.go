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
	err := json.NewDecoder(r.Body).Decode(&inputUserData)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
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

	tokenSigned, err := authPageHandlers.authService.GenerateTokens(login, user.IsAdmin, user.Version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	err = authPageHandlers.sessionService.Add(user.Email, tokenSigned, user.Version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
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
		authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
	}

	authPageHandlers.logger.Info("success login")
}

func (authPageHandlers *AuthPageHandlers) Logout(w http.ResponseWriter, r *http.Request) {
	userToken, err := r.Cookie("access")
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	tokenClaims, err := authPageHandlers.authService.IsTokenValid(userToken)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	err = authPageHandlers.sessionService.DeleteSession(tokenClaims["Login"].(string), userToken.Value)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
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

	_, err = authPageHandlers.sessionService.GetVersion(tokenClaims["Login"].(string), userToken.Value)
	if err != nil {
		authPageHandlers.logger.Info("success logout")
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

	tokenSigned, err := authPageHandlers.authService.GenerateTokens(login, false, version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	err = authPageHandlers.sessionService.Add(login, tokenSigned, version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
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
		authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
	}
}

func (authPageHandlers *AuthPageHandlers) Check(w http.ResponseWriter, r *http.Request) {
	userToken, err := r.Cookie("access")
	if err != nil {
		err = WriteError(w, myerrors.ErrNoActiveSession)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	tokenClaims, err := authPageHandlers.authService.IsTokenValid(userToken)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	err = authPageHandlers.sessionService.HasSession(tokenClaims["Login"].(string), userToken.Value)
	if err != nil {
		err = WriteError(w, myerrors.ErrNoActiveSession)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
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
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
		}
		return
	}

	err = authPageHandlers.sessionService.Add(tokenClaims["Login"].(string), tokenSigned,
		uint8(tokenClaims["Version"].(float64)))
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
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
	err = WriteSuccess(w)
	if err != nil {
		authPageHandlers.logger.Errorf("error at writing response: %v\n", err)
	}
}
