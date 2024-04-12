package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	mockService "github.com/go-park-mail-ru/2024_1_Netrunners/internal/handlers/mock"
)

func TestUserPageHandlers_GetProfileData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := mockService.NewMockAuthService(ctrl)
	mockSessionService := mockService.NewMockSessionService(ctrl)

	userHandlers := NewUserPageHandlers(mockAuthService, mockSessionService, zaptest.NewLogger(t).Sugar())

	userData := domain.User{
		Uuid: "1",
		Name: "Test User",
	}

	mockAuthService.EXPECT().GetUserDataByUuid(gomock.Any(), "1").Return(userData, nil)

	req := httptest.NewRequest("GET", "/users/1", nil)
	req = mux.SetURLVars(req, map[string]string{"uuid": "1"})
	w := httptest.NewRecorder()

	userHandlers.GetProfileData(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response profileResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.Status)
	assert.Equal(t, userData.Uuid, response.UserInfo.Uuid)
	assert.Equal(t, userData.Name, response.UserInfo.Name)
}

func TestUserPageHandlers_GetProfilePreview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := mockService.NewMockAuthService(ctrl)
	mockSessionService := mockService.NewMockSessionService(ctrl)

	userHandlers := NewUserPageHandlers(mockAuthService, mockSessionService, zaptest.NewLogger(t).Sugar())

	userData := domain.User{
		Uuid: "1",
		Name: "Test User",
	}

	mockAuthService.EXPECT().GetUserPreview(gomock.Any(), "1").Return(domain.UserPreview{
		Uuid: userData.Uuid,
		Name: userData.Name,
	}, nil)

	reqBody, err := json.Marshal(userData)
	assert.Nil(t, err)

	req := httptest.NewRequest("POST", "/profile-preview", strings.NewReader(string(reqBody)))
	w := httptest.NewRecorder()

	userHandlers.GetProfilePreview(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response profilePreviewResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.Status)
	assert.Equal(t, userData.Uuid, response.UserPreview.Uuid)
	assert.Equal(t, userData.Name, response.UserPreview.Name)
}

func TestUserPageHandlers_ProfileEditByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := mockService.NewMockAuthService(ctrl)
	mockSessionService := mockService.NewMockSessionService(ctrl)

	userHandlers := NewUserPageHandlers(mockAuthService, mockSessionService, zaptest.NewLogger(t).Sugar())

	mockUserToken := &http.Cookie{Name: "access", Value: "token"}

	mockAuthService.EXPECT().IsTokenValid(mockUserToken).Return(nil, nil)
	mockAuthService.EXPECT().GetUserDataByUuid(gomock.Any(), "1").Return(domain.User{}, nil)
	mockAuthService.EXPECT().ChangeUserPasswordByUuid(gomock.Any(), "1", "newPassword").Return(domain.User{}, nil)
	mockAuthService.EXPECT().GenerateTokens(gomock.Any(), false, gomock.Any()).Return("newToken", nil)
	mockSessionService.EXPECT().Add(gomock.Any(), "", "newToken", uint8(1)).Return(nil)

	reqBody := `{"action":"chPassword", "data":"newPassword"}`
	req := httptest.NewRequest("PUT", "/users/1", strings.NewReader(reqBody))
	req.AddCookie(mockUserToken)
	req = mux.SetURLVars(req, map[string]string{"uuid": "1"})
	w := httptest.NewRecorder()

	userHandlers.ProfileEditByUuid(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
