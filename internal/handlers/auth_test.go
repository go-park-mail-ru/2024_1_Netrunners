package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	mockService "github.com/go-park-mail-ru/2024_1_Netrunners/internal/handlers/mock"
)

func TestAuthPageHandlers_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := mockService.NewMockAuthService(ctrl)
	mockSessionService := mockService.NewMockSessionService(ctrl)

	authHandlers := NewAuthPageHandlers(mockAuthService, mockSessionService, zaptest.NewLogger(t).Sugar())

	loginData := domain.UserSignUp{
		Email:    "test@example.com",
		Password: "password",
	}

	mockAuthService.EXPECT().HasUser(gomock.Any(), loginData.Email, loginData.Password).Return(nil)
	mockAuthService.EXPECT().GetUser(gomock.Any(), loginData.Email).Return(domain.User{}, nil)
	mockAuthService.EXPECT().GenerateTokens(loginData.Email, false, gomock.Any()).Return("token", nil)
	mockSessionService.EXPECT().Add(gomock.Any(), "", "token", gomock.Any()).Return(nil)

	reqBody, err := json.Marshal(loginData)
	assert.Nil(t, err)

	req := httptest.NewRequest("POST", "/login", strings.NewReader(string(reqBody)))
	w := httptest.NewRecorder()

	authHandlers.Login(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestAuthPageHandlers_Logout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := mockService.NewMockAuthService(ctrl)
	mockSessionService := mockService.NewMockSessionService(ctrl)

	authHandlers := NewAuthPageHandlers(mockAuthService, mockSessionService, zaptest.NewLogger(t).Sugar())

	mockUserToken := &http.Cookie{Name: "access", Value: "token"}

	tokenClaims := jwt.MapClaims{"Login": "test@example.com"}

	mockAuthService.EXPECT().IsTokenValid(mockUserToken).Return(tokenClaims, nil)
	mockSessionService.EXPECT().DeleteSession(gomock.Any(), "test@example.com", "token").Return(nil)
	mockSessionService.EXPECT().GetVersion(gomock.Any(), "test@example.com", "token").Return(uint8(1),
		fmt.Errorf("err"))

	req := httptest.NewRequest("POST", "/logout", nil)
	req.AddCookie(mockUserToken)
	w := httptest.NewRecorder()

	authHandlers.Logout(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestAuthPageHandlers_Signup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := mockService.NewMockAuthService(ctrl)
	mockSessionService := mockService.NewMockSessionService(ctrl)

	authHandlers := NewAuthPageHandlers(mockAuthService, mockSessionService, zaptest.NewLogger(t).Sugar())

	signupData := domain.UserSignUp{
		Email:    "test@example.com",
		Name:     "Test User",
		Password: "password",
	}

	mockAuthService.EXPECT().CreateUser(gomock.Any(), signupData).Return(nil)
	mockAuthService.EXPECT().GenerateTokens(signupData.Email, false, gomock.Any()).Return("token", nil)
	mockSessionService.EXPECT().Add(gomock.Any(), signupData.Email, "token", gomock.Any()).Return(nil)
	mockAuthService.EXPECT().GetUser(gomock.Any(), signupData.Email).Return(domain.User{}, nil)

	reqBody, err := json.Marshal(signupData)
	assert.Nil(t, err)

	req := httptest.NewRequest("POST", "/signup", strings.NewReader(string(reqBody)))
	w := httptest.NewRecorder()

	authHandlers.Signup(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestAuthPageHandlers_Check(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := mockService.NewMockAuthService(ctrl)
	mockSessionService := mockService.NewMockSessionService(ctrl)

	authHandlers := NewAuthPageHandlers(mockAuthService, mockSessionService, zaptest.NewLogger(t).Sugar())

	mockUserToken := &http.Cookie{Name: "access", Value: "token"}

	tokenClaims := jwt.MapClaims{
		"Login":   "test@example.com",
		"IsAdmin": false,
		"Version": float64(1),
	}

	mockAuthService.EXPECT().IsTokenValid(mockUserToken).Return(tokenClaims, nil)
	mockSessionService.EXPECT().HasSession(gomock.Any(), tokenClaims["Login"].(string), mockUserToken.Value).Return(nil)
	mockAuthService.EXPECT().GenerateTokens(gomock.Any(), gomock.Any(), gomock.Any()).Return("new_token", nil)
	mockSessionService.EXPECT().Add(gomock.Any(), "test@example.com", "new_token", uint8(1)).Return(nil)

	req := httptest.NewRequest("POST", "/check", nil)
	req.AddCookie(mockUserToken)
	w := httptest.NewRecorder()

	authHandlers.Check(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
