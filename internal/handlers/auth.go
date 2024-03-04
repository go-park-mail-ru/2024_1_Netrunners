package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
)

var (
	noSuchUser                  = errors.New("no such user")
	wrongLoginOrPassword        = errors.New("wrong login or password")
	noActiveSession             = errors.New("no active session")
	notAuthorised               = errors.New("not authorised")
	tokenIsNotValid             = errors.New("token is not valid")
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
	user, err := authPageHandlers.authService.GetUser(login)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, noSuchUser)
	}

	if password != user.Password {
		WriteError(w, http.StatusInternalServerError, wrongLoginOrPassword)
	}

	accessTokenSigned, refreshTokenSigned, err := authPageHandlers.sessionService.GenerateTokens(login, user.Status, user.Version)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
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

	err = authPageHandlers.sessionService.Add(user.Login, refreshTokenSigned, user.Version)

	fmt.Println("success login")
}

func (authPageHandlers *AuthPageHandlers) Logout(w http.ResponseWriter, r *http.Request) {
	userRefreshToken, err := r.Cookie("refresh")
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
	}

	login := r.FormValue("login")

	err = authPageHandlers.sessionService.DeleteSession(login, userRefreshToken.Value)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
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
		WriteError(w, http.StatusInternalServerError, err)
	}

	accessTokenSigned, refreshTokenSigned, err := authPageHandlers.sessionService.GenerateTokens(username, status, version)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
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
		WriteError(w, http.StatusInternalServerError, err)
	}
}

func (authPageHandlers *AuthPageHandlers) Check(w http.ResponseWriter, r *http.Request) {
	userRefreshToken, err := r.Cookie("refresh")
	if err != nil {
		WriteError(w, http.StatusInternalServerError, noActiveSession)
	}

	hasSession := authPageHandlers.sessionService.HasSession(r.FormValue("login"), userRefreshToken.Value)
	if !hasSession {
		if err != nil {
			WriteError(w, http.StatusInternalServerError, noActiveSession)
		}
	}

	refreshToken, err := authPageHandlers.authService.IsTokenValid(userRefreshToken)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
	}

	if !refreshToken.Valid {
		WriteError(w, http.StatusUnauthorized, notAuthorised)
	}
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		WriteError(w, http.StatusUnauthorized, tokenIsNotValid)
	}

	login, ok := claims["Login"].(string)
	if !ok {
		WriteError(w, http.StatusUnauthorized, tokenIsNotValid)
	}
	status, ok := claims["Status"].(string)
	if !ok {
		WriteError(w, http.StatusUnauthorized, tokenIsNotValid)
	}

	ver, ok := claims["Version"].(float64)
	version := uint8(ver)
	if !ok {
		WriteError(w, http.StatusUnauthorized, tokenIsNotValid)
	}

	accessTokenSigned, refreshTokenSigned, err := authPageHandlers.sessionService.GenerateTokens(login, status, version)

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
}
