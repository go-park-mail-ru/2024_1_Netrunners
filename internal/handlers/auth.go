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

	match := ValidateLogin(login)
	if !match {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", loginIsNotValid)
		}
		return
	}

	err = ValidatePassword(password)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	err = ValidatePassword(password)
	if err != nil {
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
	username := inputUserData.Username
	password := inputUserData.Password
	fmt.Println(login)	
	match := ValidateLogin(login)
	if !match {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", loginIsNotValid)
		}
		return
	}

	fmt.Println(username)
	err = ValidateUsername(username)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	err = ValidatePassword(password)
	if err != nil {
		err = WriteError(w, err)
		if err != nil {
			fmt.Printf("error at writing response: %v\n", err)
		}
		return
	}

	
	status := "regular"
	var version uint8 = 1

	var user = domain.User{
		Login:    login,
		Username:     username,
		Password: password,
		Status:   status,
		Version:  version,
	}

	fmt.Println(user.Login)
	err = authPageHandlers.authService.CreateUser(user)
	if err != nil {
		err = WriteError(w, err)
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
	err = authPageHandlers.sessionService.CheckAllUserSessionTokens(login)
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
		fmt.Printf("error at writing response: %v\n", err)
	}
}

func (authPageHandlers *AuthPageHandlers) Check(w http.ResponseWriter, r *http.Request) {
	userRefreshToken, err := r.Cookie("refresh")
	fmt.Println(userRefreshToken)
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
	
	fmt.Println(refreshTokenClaims)	
	fmt.Println("data to check hassession:", refreshTokenClaims["Login"].(string), "|||", userRefreshToken.Value)
	hasSession := authPageHandlers.sessionService.HasSession(refreshTokenClaims["Login"].(string),
		userRefreshToken.Value)
	if hasSession {
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
	}
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
	if err != nil {	}

	err = authPageHandlers.sessionService.Add(refreshTokenClaims["Login"].(string), refreshTokenSigned,
		uint8(refreshTokenClaims["Version"].(float64)))
	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
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
		fmt.Printf("error at writing response: %v\n", err)
	}
}


func ValidateLogin(e string) bool {
    emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return emailRegex.MatchString(e)
}

func ValidateUsername(username string) error {
	if len(username) >= 4 {
		return nil
	}
	return usernameIsToShort
}

func ValidatePassword(password string) error {
	if len(password) >= 6 {
		return nil
	}
	return passwordIsToShort
}

