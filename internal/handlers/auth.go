package handlers

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
	"net/http"
)

var (
	noSuchUser            = errors.New("no such user")
	wrongLoginOrPassword  = errors.New("wrong login or password")
	tokenGenerationIssues = errors.New("token generating issues")
	noActiveSession       = errors.New("no active session")
	notAuthorised         = errors.New("not authorised")
	tokenIsNotValid       = errors.New("token is not valid")
)

type AuthPageHandlers struct {
	authService    *service.AuthService
	sessionService *service.SessionService
}

func InitAuthPageHandlers(authService *service.AuthService, sessionService *service.SessionService) *AuthPageHandlers {
	return &AuthPageHandlers{
		authService:    authService,
		sessionService: sessionService,
	}
}

func (authPageHandlers *AuthPageHandlers) Login(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	user, err := authPageHandlers.authService.GetUser(login)

	if err != nil {
		// нет такого пользователя в базе
		errs := WriteError(w, 500, noSuchUser)
		if errs != nil {
			return
		}
		return
	}

	if password != user.Password {
		// неверный логин или пароль
		errs := WriteError(w, 500, wrongLoginOrPassword)
		if errs != nil {
			return
		}
		return
	}

	accessTokenSigned, refreshTokenSigned, err := authPageHandlers.sessionService.GenerateTokens(login, user.Status, user.Version)
	if err != nil {
		// проблемы с генерацией токена
		errs := WriteError(w, 500, err)
		if errs != nil {
			return
		}
		return
	}

	accessCookie := &http.Cookie{
		Name:     "access",
		Value:    accessTokenSigned,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   5 * 60, // 5 mins
	}

	refreshCookie := &http.Cookie{
		Name:     "refresh",
		Value:    refreshTokenSigned,
		Path:     "/auth",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   48 * 3600, // 48 hours
	}

	http.SetCookie(w, refreshCookie)
	http.SetCookie(w, accessCookie)

	err = authPageHandlers.sessionService.Add(user.Login, refreshTokenSigned, user.Version)

	fmt.Println("success login")
}

func (authPageHandlers *AuthPageHandlers) Logout(w http.ResponseWriter, r *http.Request) {
	userRefreshToken, err := r.Cookie("refresh")
	if err != nil {
		errs := WriteError(w, 500, err)
		if errs != nil {
			return
		}
		return
	}
	login := r.FormValue("login")

	if err != nil {
		errs := WriteError(w, 500, err)
		if errs != nil {
			return
		}
		return
	}

	err = authPageHandlers.sessionService.DeleteSession(login, userRefreshToken.Value)

	if err != nil {
		errs := WriteError(w, 500, err)
		if errs != nil {
			return
		}
		return
	}

	refreshCookie := &http.Cookie{
		Name:   "refresh",
		MaxAge: 0,
	}
	accessCookie := &http.Cookie{
		Name:   "access",
		MaxAge: 0,
	}

	http.SetCookie(w, refreshCookie)
	http.SetCookie(w, accessCookie)

	_, err = authPageHandlers.sessionService.GetVersion(login, userRefreshToken.Value)
	if err != nil {
		fmt.Println("success logout")
	}
}

func (authPageHandlers *AuthPageHandlers) Signup(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	username := r.FormValue("username")
	password := r.FormValue("password")

	status := "regular"
	var version uint8 = 1

	var user = domain.User{
		Login:    login,
		Name:     username,
		Password: password,
		Status:   status,
		Version:  version,
	}

	err := authPageHandlers.authService.Create(user)
	if err != nil {
		errs := WriteError(w, 500, err)
		if errs != nil {
			return
		}
		return
	}

	accessTokenSigned, refreshTokenSigned, err := authPageHandlers.sessionService.GenerateTokens(username, status, version)
	if err != nil {
		errs := WriteError(w, 500, err)
		if errs != nil {
			return
		}
		return
	}

	if err != nil {
		errs := WriteError(w, 500, err)
		if errs != nil {
			return
		}
		return
	}

	accessCookie := &http.Cookie{
		Name:     "access",
		Value:    accessTokenSigned,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   5 * 60, // 5 mins
	}

	refreshCookie := &http.Cookie{
		Name:     "refresh",
		Value:    refreshTokenSigned,
		Path:     "/auth",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   48 * 3600, // 48 hours
	}

	http.SetCookie(w, refreshCookie)
	http.SetCookie(w, accessCookie)

	err = authPageHandlers.sessionService.Add(login, refreshTokenSigned, version)

	if err != nil {
		errs := WriteError(w, 500, err)
		if errs != nil {
			return
		}
		return
	}

	fmt.Println("success signup")
}

func (authPageHandlers *AuthPageHandlers) Check(w http.ResponseWriter, r *http.Request) {
	userRefreshToken, err := r.Cookie("refresh")
	if err != nil {
		// у юзера нет активной сессии
		errs := WriteError(w, 500, noActiveSession)
		if errs != nil {
			return
		}
		return
	}

	hasSession := authPageHandlers.sessionService.HasSession(r.FormValue("login"), userRefreshToken.Value)

	if !hasSession {
		if err != nil {
			// нет такой сессии на сервере
			errs := WriteError(w, 500, noActiveSession)
			if errs != nil {
				return
			}
			return
		}
	}

	refreshToken, err := authPageHandlers.authService.IsTokenValid(userRefreshToken)

	if err != nil {
		// не удалось распарсить токен
		errs := WriteError(w, 500, err)
		if errs != nil {
			return
		}
		return
	}

	if !refreshToken.Valid {
		/// не авторизован
		errs := WriteError(w, 401, notAuthorised)
		if errs != nil {
			return
		}
		return
	}
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		errs := WriteError(w, 401, tokenIsNotValid)
		if errs != nil {
			return
		}
		return
	}

	login, ok := claims["Login"].(string)
	if !ok {

		errs := WriteError(w, 401, tokenIsNotValid)
		if errs != nil {
			return
		}
		return
	}
	status, ok := claims["Status"].(string)
	if !ok {

		errs := WriteError(w, 401, tokenIsNotValid)
		if errs != nil {
			return
		}
		return
	}

	ver, ok := claims["Version"].(float64)
	version := uint8(ver)

	if !ok {
		errs := WriteError(w, 401, tokenIsNotValid)
		if errs != nil {
			return
		}
		return
	}

	accessTokenSigned, refreshTokenSigned, err := authPageHandlers.sessionService.GenerateTokens(login, status, version)

	accessCookie := &http.Cookie{
		Name:     "access",
		Value:    accessTokenSigned,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   5 * 60, // 5 mins
	}

	refreshCookie := &http.Cookie{
		Name:     "refresh",
		Value:    refreshTokenSigned,
		Path:     "/auth",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   48 * 3600, // 48 hours
	}

	http.SetCookie(w, refreshCookie)
	http.SetCookie(w, accessCookie)

	err = authPageHandlers.sessionService.Add(login, refreshTokenSigned, version)
}
