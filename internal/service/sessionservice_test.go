package service

import (
	"testing"

	"go.uber.org/zap"

	mockService "github.com/go-park-mail-ru/2024_1_Netrunners/internal/service/mock"
	"github.com/golang/mock/gomock"
)

func TestAddSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMocksessionStorage(ctrl)
	mockLogger := zap.NewExample().Sugar()

	service := NewSessionService(mockStorage, mockLogger)

	login := "testuser"
	token := "token123"
	version := uint8(1)

	mockStorage.EXPECT().Add(login, token, version).Return(nil)

	err := service.Add(login, token, "1", version)

	if err != nil {
		t.Errorf("AddSession returned an unexpected error: %v", err)
	}
}

func TestDeleteSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMocksessionStorage(ctrl)
	mockLogger := zap.NewExample().Sugar()

	service := NewSessionService(mockStorage, mockLogger)

	login := "testuser"
	token := "token123"

	mockStorage.EXPECT().DeleteSession(login, token).Return(nil)

	err := service.DeleteSession(login, token, "1")

	if err != nil {
		t.Errorf("DeleteSession returned an unexpected error: %v", err)
	}
}

func TestUpdateSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMocksessionStorage(ctrl)
	mockLogger := zap.NewExample().Sugar()

	service := NewSessionService(mockStorage, mockLogger)

	login := "testuser"
	token := "token123"

	mockStorage.EXPECT().Update(login, token).Return(nil)

	err := service.Update(login, token, "1")

	if err != nil {
		t.Errorf("UpdateSession returned an unexpected error: %v", err)
	}
}

func TestCheckVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMocksessionStorage(ctrl)
	mockLogger := zap.NewExample().Sugar()

	service := NewSessionService(mockStorage, mockLogger)

	login := "testuser"
	token := "token123"
	usersVersion := uint8(2)

	mockStorage.EXPECT().CheckVersion(login, token, usersVersion).Return(true, nil)

	hasSession, err := service.CheckVersion(login, token, "1", usersVersion)

	if err != nil {
		t.Errorf("CheckVersion returned an unexpected error: %v", err)
	}

	if !hasSession {
		t.Error("CheckVersion returned unexpected result: expected true, got false")
	}
}

func TestGetVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMocksessionStorage(ctrl)
	mockLogger := zap.NewExample().Sugar()

	service := NewSessionService(mockStorage, mockLogger)

	login := "testuser"
	token := "token123"
	expectedVersion := uint8(3)

	mockStorage.EXPECT().GetVersion(login, token).Return(expectedVersion, nil)

	version, err := service.GetVersion(login, token, "1")

	if err != nil {
		t.Errorf("GetVersion returned an unexpected error: %v", err)
	}

	if version != expectedVersion {
		t.Errorf("GetVersion returned unexpected version: expected %d, got %d", expectedVersion, version)
	}
}

func TestHasSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMocksessionStorage(ctrl)
	mockLogger := zap.NewExample().Sugar()

	service := NewSessionService(mockStorage, mockLogger)

	login := "testuser"
	token := "token123"

	mockStorage.EXPECT().HasSession(login, token).Return(nil)

	err := service.HasSession(login, token, "1")

	if err != nil {
		t.Errorf("HasSession returned an unexpected error: %v", err)
	}
}

func TestCheckAllUserSessionTokens(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMocksessionStorage(ctrl)
	mockLogger := zap.NewExample().Sugar()

	service := NewSessionService(mockStorage, mockLogger)

	login := "testuser"

	mockStorage.EXPECT().CheckAllUserSessionTokens(login).Return(nil)

	err := service.CheckAllUserSessionTokens(login, "1")

	if err != nil {
		t.Errorf("CheckAllUserSessionTokens returned an unexpected error: %v", err)
	}
}
