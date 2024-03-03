package handlers

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/repository/cache"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
	"net/http"
)

var (
	notPOSTMethod         = errors.New("метод передачи не POST")
	SECRET                = []byte("SECRETKEY")
	noSuchUser            = errors.New("нет такого пользователя")
	wrongLoginOrPassword  = errors.New("неверный логин или пароль")
	tokenGenerationIssues = errors.New("проблемы с генерацией токена")
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
	login := r.FormValue("login")
	password := r.FormValue("password")
	user, err := authPageHandlers.authService.GetUser(login)
	if err != nil {
		// нет такого пользователя в базе
		errs := WriteError(w, 500, noSuchUser)
		if errs != nil {
			return
		}
	}

	if password != user.Password {
		// неверный логин или пароль
		return
	}

	accessTokenSigned, refreshTokenSigned, err := service.GenerateTokens(login, user.Status, user.Version)
	if err != nil {
		// проблемы с генерацией токена
		return
	}

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
	login := r.FormValue("login")

	err = authPageHandlers.cache.Delete(login, userRefreshToken.Value)
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

func (authPageHandlers *AuthPageHandlers) Signup(w http.ResponseWriter, r *http.Request) {

	login := r.FormValue("login")
	username := r.FormValue("username")
	password := r.FormValue("password")
	status := "regular"
	version := 1

	var user = domain.User{
		Login:    login,
		Name:     username,
		Password: password,
		Status:   status,
		Version:  version,
	}

	err := authPageHandlers.authService.Create(user)
	if err != nil {
		return
	}

	accessTokenSigned, refreshTokenSigned, err := service.GenerateTokens(username, status, version)
	if err != nil {
		return
	}

	err = authPageHandlers.cache.Add(login, refreshTokenSigned, version)
	if err != nil {
		return
	}

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
		MaxAge:   5 * 60, // 5 mins
	}

	http.SetCookie(w, refreshCookie)
	http.SetCookie(w, accessCookie)
}

// если кука есть, проверяем на валидность, если не валидная, проверяем рефрешТокен на валидность, если валиден, то
// 		создаем новый рефреш и аксес, и отправляем их
// если кука есть, проверяем на валидность, если не валидная, проверяем рефрешТокен на валидность, если невалиден,
//		то возвращаем 401

func (authPageHandlers *AuthPageHandlers) Check(w http.ResponseWriter, r *http.Request) {
	userRefreshToken, err := r.Cookie("refresh")
	if err != nil {
		// у юзера нет активной сессии
		return
	}

	refreshtoken, err := jwt.Parse(userRefreshToken.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return SECRET, nil
	})
	if err != nil {
		// не удалось распарсить токен
		return
	}

	if !refreshtoken.Valid {
		/// не авторизован
		return
	}

	claims, ok := refreshtoken.Claims.(jwt.MapClaims)
	if !ok {
		return
	}

	login, ok := claims["login"].(string)
	if !ok {
		return
	}
	status, ok := claims["status"].(string)
	if !ok {
		return
	}
	version, ok := claims["version"].(int)
	if !ok {
		return
	}

	accessTokenSigned, refreshTokenSigned, err := service.GenerateTokens(login, status, version)

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
