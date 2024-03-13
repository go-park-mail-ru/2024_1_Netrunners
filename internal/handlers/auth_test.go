package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	mycache "github.com/go-park-mail-ru/2024_1_Netrunners/internal/repository/cache"
	mockdb "github.com/go-park-mail-ru/2024_1_Netrunners/internal/repository/mockDB"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAuthPageHandlers_Signup(t *testing.T) {
	validCases := []struct {
		testName string
		json     []byte
		Login    string
		Name     string
		Password string
	}{
		{
			"valid json",
			[]byte(`{"login": "cakethefake@mail.ru", "username": "Danya", "password": "12345678"}`),
			"cakethefake@mail.ru",
			"Danya",
			"12345678",
		},
	}

	cacheStorage := mycache.InitSessionStorage()
	authStorage := mockdb.InitUsersMockDB()

	sessionService := service.InitSessionService(cacheStorage)
	authService := service.InitAuthService(authStorage)

	authPageHandlers := InitAuthPageHandlers(authService, sessionService)
	handler := http.HandlerFunc(authPageHandlers.Signup)

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(currentCase.json))
			handler.ServeHTTP(rec, req)

			err := authPageHandlers.authService.HasUser(currentCase.Login, currentCase.Password)
			if err != nil {
				t.Error(err)
			}

			refreshToken := ""
			cookies := rec.Result().Cookies()
			for _, cookie := range cookies {
				switch {
				case cookie.Name == "refresh":
					refreshToken = cookie.Value
				case cookie.Name != "access":
					err = fmt.Errorf("no cookie")
				}
			}
			if err != nil {
				t.Error(err)
			}

			err = authPageHandlers.sessionService.HasSession(currentCase.Login, refreshToken)
			if err != nil {
				t.Error(err)
			}

			if rec.Result().StatusCode != http.StatusOK {
				t.Error(fmt.Errorf("transport status code is not 200"))
			}

			data, err := io.ReadAll(rec.Result().Body)
			if err != nil {
				t.Error(err)
			}

			var responseData map[string]interface{}
			err = json.Unmarshal(data, &responseData)
			if err != nil {
				t.Error(err)
			}

			if responseData["status"].(float64) != 200 {
				t.Error(fmt.Errorf("server status code is not 200"))
			}
		})
	}

	invalidCases := []struct {
		testName string
		json     []byte
	}{
		{
			"login is too short",
			[]byte(`{"login": "a@a.a", "username": "Danya", "password": "12345678"}`),
		},
		{
			"login is not an email",
			[]byte(`{"login": "Danya", "username": "Danya", "password": "12345678"}`),
		},
		{
			"username is too short",
			[]byte(`{"login": "cakethefake@mail.ru", "username": "A?", "password": "12345678"}`),
		},
		{
			"password is too short",
			[]byte(`{"login": "cakethefake@mail.ru", "username": "Danya", "password": "777"}`),
		},
		{
			"user already exists",
			[]byte(`{"login": "cakethefake@mail.ru", "username": "Danya", "password": "12345678"}`),
		},
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(currentCase.json))
			handler.ServeHTTP(rec, req)

			data, err := io.ReadAll(rec.Result().Body)
			if err != nil {
				t.Error(err)
			}

			var responseData map[string]interface{}
			err = json.Unmarshal(data, &responseData)
			if err != nil {
				t.Error(err)
			}

			if responseData["status"].(float64) == 200 {
				t.Error(fmt.Errorf("server status code is 200"))
			}
		})
	}
}

func TestAuthPageHandlers_Login(t *testing.T) {
	validCases := []struct {
		testName string
		json     []byte
		Login    string
		Password string
	}{
		{
			"valid json",
			[]byte(`{"login": "cakethefake@mail.ru", "password": "12345678"}`),
			"cakethefake@mail.ru",
			"12345678",
		},
	}

	cacheStorage := mycache.InitSessionStorage()
	authStorage := mockdb.InitUsersMockDB()

	sessionService := service.InitSessionService(cacheStorage)
	authService := service.InitAuthService(authStorage)

	authPageHandlers := InitAuthPageHandlers(authService, sessionService)
	handlerLogin := http.HandlerFunc(authPageHandlers.Login)

	user := domain.User{
		Login:    "cakethefake@mail.ru",
		Name:     "Danya",
		Password: "12345678",
	}

	err := authPageHandlers.authService.CreateUser(user)
	if err != nil {
		t.Error(err)
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {

			rec := httptest.NewRecorder()
			reqLogin, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(currentCase.json))
			handlerLogin.ServeHTTP(rec, reqLogin)

			refreshToken := ""
			cookies := rec.Result().Cookies()
			for _, cookie := range cookies {
				switch {
				case cookie.Name == "refresh":
					refreshToken = cookie.Value
				case cookie.Name != "access":
					err = fmt.Errorf("no cookie")
				}
			}
			if err != nil {
				t.Error(err)
			}

			err = authPageHandlers.sessionService.HasSession(currentCase.Login, refreshToken)
			if err != nil {
				t.Error(err)
			}

			if rec.Result().StatusCode != http.StatusOK {
				t.Error(fmt.Errorf("transport status code is not 200"))
			}

			data, err := io.ReadAll(rec.Result().Body)
			if err != nil {
				t.Error(err)
			}

			var responseData map[string]interface{}
			err = json.Unmarshal(data, &responseData)
			if err != nil {
				t.Error(err)
			}

			if responseData["status"].(float64) != 200 {
				t.Error(fmt.Errorf("server status code is not 200"))
			}
		})
	}

	invalidCases := []struct {
		testName string
		json     []byte
	}{
		{
			"no such user found",
			[]byte(`{"login": "a@a.a", "password": "12345678"}`),
		},
		{
			"wrong password",
			[]byte(`{"login": "cakethefake@mail.ru", "password": "12345677"}`),
		},
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {

			rec := httptest.NewRecorder()
			reqLogin, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(currentCase.json))
			handlerLogin.ServeHTTP(rec, reqLogin)

			err := fmt.Errorf("")

			data, err := io.ReadAll(rec.Result().Body)
			if err != nil {
				t.Error(err)
			}

			var responseData map[string]interface{}
			err = json.Unmarshal(data, &responseData)
			if err != nil {
				t.Error(err)
			}

			if responseData["status"].(float64) == 200 {
				t.Error(fmt.Errorf("server status code is 200"))
			}
		})
	}
}

func TestAuthPageHandlers_Logout(t *testing.T) {
	cacheStorage := mycache.InitSessionStorage()
	authStorage := mockdb.InitUsersMockDB()

	sessionService := service.InitSessionService(cacheStorage)
	authService := service.InitAuthService(authStorage)

	authPageHandlers := InitAuthPageHandlers(authService, sessionService)

	user := domain.User{
		Login:    "cakethefake@mail.ru",
		Name:     "Danya",
		Password: "12345678",
	}

	err := authPageHandlers.authService.CreateUser(user)
	if err != nil {
		t.Error(err)
	}

	recLogin := httptest.NewRecorder()
	reqLogin, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer([]byte(
		`{"login": "cakethefake@mail.ru", 
		 "password": "12345678"}`)))
	handlerLogin := http.HandlerFunc(authPageHandlers.Login)
	handlerLogin.ServeHTTP(recLogin, reqLogin)

	var refreshToken string
	cookiesLogin := recLogin.Result().Cookies()
	for _, cookie := range cookiesLogin {
		if cookie.Name == "refresh" {
			refreshToken = cookie.Value
		}
	}

	recLogout := httptest.NewRecorder()
	reqLogout, _ := http.NewRequest("POST", "/auth/logout", nil)
	reqLogout.AddCookie(&http.Cookie{Name: "refresh", Value: refreshToken})
	handlerLogout := http.HandlerFunc(authPageHandlers.Logout)
	handlerLogout.ServeHTTP(recLogout, reqLogout)

	err = authPageHandlers.sessionService.HasSession("cakethefake@mail.ru", refreshToken)
	if err == nil {
		t.Errorf("Expected session to be deleted, but it still exists")
	}

	cookiesLogout := recLogout.Result().Cookies()
	for _, cookie := range cookiesLogout {
		if cookie.Name == "access" || cookie.Name == "refresh" {
			if cookie.Value != "" {
				t.Errorf("Expected cookie value to be empty, got: %s", cookie.Value)
			}
		}
	}

	if recLogout.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got: %d", recLogout.Result().StatusCode)
	}
}

func TestAuthPageHandlers_Check(t *testing.T) {
	cacheStorage := mycache.InitSessionStorage()
	authStorage := mockdb.InitUsersMockDB()

	sessionService := service.InitSessionService(cacheStorage)
	authService := service.InitAuthService(authStorage)

	authPageHandlers := InitAuthPageHandlers(authService, sessionService)

	user := domain.User{
		Login:    "cakethefake@mail.ru",
		Name:     "Danya",
		Password: "12345678",
	}

	err := authPageHandlers.authService.CreateUser(user)
	if err != nil {
		t.Error(err)
	}

	validCases := []struct {
		testName string
		json     []byte
		Login    string
		Password string
	}{
		{
			"valid json",
			[]byte(`{"login": "cakethefake@mail.ru", "password": "12345678"}`),
			"cakethefake@mail.ru",
			"12345678",
		},
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			recLogin := httptest.NewRecorder()
			reqLogin, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(currentCase.json))
			handlerLogin := http.HandlerFunc(authPageHandlers.Login)
			handlerLogin.ServeHTTP(recLogin, reqLogin)
			time.Sleep(time.Second * 1)

			var refreshToken string
			cookiesLogin := recLogin.Result().Cookies()
			for _, cookie := range cookiesLogin {
				if cookie.Name == "refresh" {
					refreshToken = cookie.Value
				}
			}

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/auth/check", nil)
			req.AddCookie(&http.Cookie{Name: "refresh", Value: refreshToken})
			handlerCheck := http.HandlerFunc(authPageHandlers.Check)
			handlerCheck.ServeHTTP(rec, req)

			data, err := io.ReadAll(rec.Result().Body)
			if err != nil {
				t.Error(err)
			}

			var responseData map[string]interface{}
			err = json.Unmarshal(data, &responseData)
			if err != nil {
				t.Error(err)
			}

			if responseData["status"].(float64) != 200 {
				t.Error(fmt.Errorf("server status code is not 200"))
			}
		})
	}

	invelidCase1 := []struct {
		testName string
		json     []byte
		Login    string
		Password string
	}{
		{
			"no such session",
			[]byte(`{"login": "cakethefake@mail.ru", "password": "12345678"}`),
			"cakethefake@mail.ru",
			"12345678",
		},
	}

	for _, currentCase := range invelidCase1 {
		t.Run(currentCase.testName, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/auth/check", nil)
			handlerCheck := http.HandlerFunc(authPageHandlers.Check)
			handlerCheck.ServeHTTP(rec, req)

			data, err := io.ReadAll(rec.Result().Body)
			if err != nil {
				t.Error(err)
			}

			var responseData map[string]interface{}
			err = json.Unmarshal(data, &responseData)
			if err != nil {
				t.Error(err)
			}

			if responseData["status"].(float64) == 200 {
				t.Error(fmt.Errorf("server status code is not 200"))
			}
		})
	}

	invelidCase2 := []struct {
		testName string
		json     []byte
		Login    string
		Password string
	}{
		{
			"invalid refresh token in cookie",
			[]byte(`{"login": "cakethefake@mail.ru", "password": "12345678"}`),
			"cakethefake@mail.ru",
			"12345678",
		},
	}

	for _, currentCase := range invelidCase2 {
		t.Run(currentCase.testName, func(t *testing.T) {
			recLogin := httptest.NewRecorder()
			reqLogin, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(currentCase.json))
			handlerLogin := http.HandlerFunc(authPageHandlers.Login)
			handlerLogin.ServeHTTP(recLogin, reqLogin)
			time.Sleep(time.Second * 1)

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/auth/check", nil)
			req.AddCookie(&http.Cookie{Name: "refresh", Value: "refreshToken"})
			handlerCheck := http.HandlerFunc(authPageHandlers.Check)
			handlerCheck.ServeHTTP(rec, req)

			data, err := io.ReadAll(rec.Result().Body)
			if err != nil {
				t.Error(err)
			}

			var responseData map[string]interface{}
			err = json.Unmarshal(data, &responseData)
			if err != nil {
				t.Error(err)
			}

			if responseData["status"].(float64) == 200 {
				t.Error(fmt.Errorf("server status code is not 200"))
			}
		})
	}
}
