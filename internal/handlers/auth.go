package handlers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"

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
	err := authPageHandlers.authService.HasUser(login, password)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	user, _ := authPageHandlers.authService.GetUser(login)

	accessTokenSigned, refreshTokenSigned, err := authPageHandlers.sessionService.GenerateTokens(login, user.Status, user.Version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	err = authPageHandlers.sessionService.Add(user.Login, refreshTokenSigned, user.Version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	accessCookie := &http.Cookie{
		Name:     "access",
		Value:    accessTokenSigned,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   accessCookieExpirationTime,
	}

	refreshCookie := &http.Cookie{
		Name:     "refresh",
		Value:    refreshTokenSigned,
		Path:     "/auth",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   refreshCookieExpirationTime,
	}

	http.SetCookie(w, refreshCookie)
	http.SetCookie(w, accessCookie)
	err = WriteSuccess(w)
	if err != nil {
		fmt.Printf("error at writing response: %v\n", err)
	}

	fmt.Println("success login")
}

func (authPageHandlers *AuthPageHandlers) Logout(w http.ResponseWriter, r *http.Request) {
	userRefreshToken, err := r.Cookie("refresh")
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	refreshToken, err := authPageHandlers.authService.IsTokenValid(userRefreshToken)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	if !refreshToken.Valid {
		err = WriteError(w, myerrors.ErrNotAuthorised)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		err = WriteError(w, myerrors.ErrTokenIsNotValid)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	login, ok := claims["Login"].(string)
	if !ok {
		err = WriteError(w, myerrors.ErrTokenIsNotValid)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	err = authPageHandlers.sessionService.DeleteSession(login, userRefreshToken.Value)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	accessCookie := &http.Cookie{
		Name:     "access",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   0,
	}

	refreshCookie := &http.Cookie{
		Name:     "refresh",
		Value:    "",
		Path:     "/auth",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   0,
	}

	http.SetCookie(w, refreshCookie)
	http.SetCookie(w, accessCookie)

	_, err = authPageHandlers.sessionService.GetVersion(login, userRefreshToken.Value)
	if err != nil {
		fmt.Println("success logout")
	}

	err = WriteSuccess(w)
	if err != nil {
		fmt.Printf("error at writing response: %v\n", err)
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

	err := authPageHandlers.authService.CreateUser(user)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	accessTokenSigned, refreshTokenSigned, err := authPageHandlers.sessionService.GenerateTokens(username, status, version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	accessCookie := &http.Cookie{
		Name:     "access",
		Value:    accessTokenSigned,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   accessCookieExpirationTime,
	}

	refreshCookie := &http.Cookie{
		Name:     "refresh",
		Value:    refreshTokenSigned,
		Path:     "/auth",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   refreshCookieExpirationTime,
	}

	http.SetCookie(w, refreshCookie)
	http.SetCookie(w, accessCookie)

	err = authPageHandlers.sessionService.Add(login, refreshTokenSigned, version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	err = WriteSuccess(w)
	if err != nil {
		fmt.Printf("error at writing response: %v\n", err)
	}
}

func (authPageHandlers *AuthPageHandlers) Check(w http.ResponseWriter, r *http.Request) {
	userRefreshToken, err := r.Cookie("refresh")
	if err != nil {
		err = WriteError(w, myerrors.ErrNoActiveSession)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	hasSession := authPageHandlers.sessionService.HasSession(r.FormValue("login"), userRefreshToken.Value)
	if !hasSession {
		if err != nil {
			err = WriteError(w, myerrors.ErrNoActiveSession)
			if err != nil {
				fmt.Printf("error at writing response: %v\n", err)
			}
		}
		return
	}

	refreshToken, err := authPageHandlers.authService.IsTokenValid(userRefreshToken)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	if !refreshToken.Valid {
		err = WriteError(w, myerrors.ErrNotAuthorised)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		err = WriteError(w, myerrors.ErrTokenIsNotValid)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	login, ok := claims["Login"].(string)
	if !ok {
		err = WriteError(w, myerrors.ErrTokenIsNotValid)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}
	status, ok := claims["Status"].(string)
	if !ok {
		err = WriteError(w, myerrors.ErrTokenIsNotValid)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	ver, ok := claims["Version"].(float64)
	version := uint8(ver)
	if !ok {
		err = WriteError(w, myerrors.ErrTokenIsNotValid)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	accessTokenSigned, refreshTokenSigned, err := authPageHandlers.sessionService.GenerateTokens(login, status, version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	err = authPageHandlers.sessionService.Add(login, refreshTokenSigned, version)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	accessCookie := &http.Cookie{
		Name:     "access",
		Value:    accessTokenSigned,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   accessCookieExpirationTime,
	}

	refreshCookie := &http.Cookie{
		Name:     "refresh",
		Value:    refreshTokenSigned,
		Path:     "/auth",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   refreshCookieExpirationTime,
	}

	http.SetCookie(w, refreshCookie)
	http.SetCookie(w, accessCookie)
	err = WriteSuccess(w)
	if err != nil {
		fmt.Printf("error at writing response: %v\n", err)
	}
}
