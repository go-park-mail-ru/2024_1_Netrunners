package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
)

var (
	loginIsNotValid             = errors.New("login is not valid")
	passwordIsToShort           = errors.New("password is too short")
	usernameIsToShort           = errors.New("username is too short")
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
	var inputUserData domain.User
	err := json.NewDecoder(r.Body).Decode(&inputUserData)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}
	login := inputUserData.Login
	password := inputUserData.Password

	match, err := ValidateLogin(login)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}
	if !match {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", loginIsNotValid)
		}
		return
	}

	match = ValidatePassword(password)
	if !match {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", passwordIsToShort)
		}
		return
	}

	err = authPageHandlers.authService.HasUser(login, password)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	user, err := authPageHandlers.authService.GetUser(login)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	accessTokenSigned, refreshTokenSigned, err := authPageHandlers.sessionService.GenerateTokens(login,
		user.Status, user.Version)
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

	refreshTokenClaims, err := authPageHandlers.authService.IsTokenValid(userRefreshToken)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	err = authPageHandlers.sessionService.DeleteSession(refreshTokenClaims["Login"].(string), userRefreshToken.Value)
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

	_, err = authPageHandlers.sessionService.GetVersion(refreshTokenClaims["Login"].(string), userRefreshToken.Value)
	if err != nil {
		fmt.Println("success logout")
	}

	err = WriteSuccess(w)
	if err != nil {
		fmt.Printf("error at writing response: %v\n", err)
	}
}

func (authPageHandlers *AuthPageHandlers) Signup(w http.ResponseWriter, r *http.Request) {
	var inputUserData domain.User
	err := json.NewDecoder(r.Body).Decode(&inputUserData)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}
	login := inputUserData.Login
	username := inputUserData.Name
	password := inputUserData.Password

	match, err := ValidateLogin(login)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}
	if !match {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", loginIsNotValid)
		}
		return
	}

	match = ValidateUsername(username)
	if !match {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", usernameIsToShort)
		}
		return
	}

	match = ValidatePassword(password)
	if !match {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", passwordIsToShort)
		}
		return
	}

	status := "regular"
	var version uint8 = 1

	var user = domain.User{
		Login:    login,
		Name:     username,
		Password: password,
		Status:   status,
		Version:  version,
	}

	err = authPageHandlers.authService.CreateUser(user)
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

func (authPageHandlers *AuthPageHandlers) Check(w http.ResponseWriter, r *http.Request) {
	userRefreshToken, err := r.Cookie("refresh")
	if err != nil {
		err = WriteError(w, myerrors.ErrNoActiveSession)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	refreshTokenClaims, err := authPageHandlers.authService.IsTokenValid(userRefreshToken)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	hasSession := authPageHandlers.sessionService.HasSession(refreshTokenClaims["Login"].(string),
		userRefreshToken.Value)
	if !hasSession {
		err = WriteError(w, myerrors.ErrNoActiveSession)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	userAccessToken, err := r.Cookie("access")
	if err == nil {
		_, err = authPageHandlers.authService.IsTokenValid(userAccessToken)
		if err == nil {
			err = WriteSuccess(w)
			if err != nil {
				fmt.Printf("error at writing response: %v\n", err)
			}
			return
		}
	}

	accessTokenSigned, refreshTokenSigned, err := authPageHandlers.sessionService.GenerateTokens(
		refreshTokenClaims["Login"].(string), refreshTokenClaims["Status"].(string),
		uint8(refreshTokenClaims["Version"].(float64)))
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	err = authPageHandlers.sessionService.Add(
		refreshTokenClaims["Login"].(string), refreshTokenSigned,
		uint8(refreshTokenClaims["Version"].(float64)))
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

func ValidateLogin(login string) (match bool, err error) {
	match, err = regexp.MatchString(`/^[\w-.]+@([\w-]+\.)+[\w-]{2,4}$/`, login)
	return match, err
}

func ValidateUsername(username string) bool {
	if len(username) >= 4 {
		return true
	}
	return false
}

func ValidatePassword(password string) bool {
	if len(password) >= 6 {
		return true
	}
	return false
}
