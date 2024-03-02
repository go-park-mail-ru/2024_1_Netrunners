package handlers

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/repository/cache"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
	"net/http"
)

var (
	notPOSTMethod = errors.New("метод передачи не POST")
	SECRET        = []byte("SECRETKEY")
)

type AuthPageHandlers struct {
	authService *service.AuthService
	cache       *cache.SessionStorage
}

func InitAuthPageHandlers(authService *service.AuthService, cache *cache.SessionStorage) *AuthPageHandlers {
	return &AuthPageHandlers{
		authService: authService,
		cache:       cache,
	}
}

func (authPageHandlers *AuthPageHandlers) Login(w http.ResponseWriter, r *http.Request) {
	user, err := authPageHandlers.authService.GetUser(r.FormValue("username"))
	if err != nil {
		return
	}

	username := user.Login
	status := user.Status

	accessTokenSigned, refreshTokenSigned, err := service.GenerateTokens("", username, status)

	refreshCookie := &http.Cookie{
		Name:     "refresh",
		Value:    accessTokenSigned,
		Path:     "/auth",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   48 * 3600, // 48 hours
	}

	accessCookie := &http.Cookie{
		Name:     "access",
		Value:    refreshTokenSigned,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   300, // 5 mins
	}
	http.SetCookie(w, refreshCookie)
	http.SetCookie(w, accessCookie)
}

func (authPageHandlers *AuthPageHandlers) Logout(w http.ResponseWriter, r *http.Request) {

	userRefreshToken, err := r.Cookie("refresh")
	if err != nil {
		return
	}

	token, err := jwt.Parse(userRefreshToken.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return SECRET, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return
	}

	login, ok := claims["login"].(string)
	if !ok {
		return
	}

	err = authPageHandlers.cache.Delete(login)
	if err != nil {
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

}

func (mainPageHandlers *AuthPageHandlers) Signup(w http.ResponseWriter, r *http.Request) {

}

func (mainPageHandlers *AuthPageHandlers) Check(w http.ResponseWriter, r *http.Request) {

}
