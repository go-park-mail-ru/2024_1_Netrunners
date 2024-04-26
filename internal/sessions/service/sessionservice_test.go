package service

import (
	"context"
	"errors"
	"testing"

	"go.uber.org/zap"

	mockService "github.com/go-park-mail-ru/2024_1_Netrunners/internal/sessions/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAddSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMocksessionStorage(ctrl)
	mockLogger := zap.NewExample().Sugar()

	service := NewSessionService(mockStorage, mockLogger)

	login := "testuser"
	token := "token123"
	version := uint32(1)

	mockStorage.EXPECT().Add(login, token, version).Return(nil)

	err := service.Add(context.Background(), login, token, version)

	if err != nil {
		t.Errorf("AddSession returned an unexpected error: %v", err)
	}

	mockStorage.EXPECT().Add(login, token, version).Return(errors.New(""))

	err = service.Add(context.Background(), login, token, version)

	assert.Error(t, err)
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

	err := service.DeleteSession(context.Background(), login, token)

	if err != nil {
		t.Errorf("DeleteSession returned an unexpected error: %v", err)
	}

	mockStorage.EXPECT().DeleteSession(login, token).Return(errors.New(""))

	err = service.DeleteSession(context.Background(), login, token)

	assert.Error(t, err)
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

	err := service.Update(context.Background(), login, token)

	if err != nil {
		t.Errorf("UpdateSession returned an unexpected error: %v", err)
	}

	mockStorage.EXPECT().Update(login, token).Return(errors.New(""))

	err = service.Update(context.Background(), login, token)

	assert.Error(t, err)
}

func TestCheckVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMocksessionStorage(ctrl)
	mockLogger := zap.NewExample().Sugar()

	service := NewSessionService(mockStorage, mockLogger)

	login := "testuser"
	token := "token123"
	usersVersion := uint32(2)

	mockStorage.EXPECT().CheckVersion(login, token, usersVersion).Return(true, nil)

	hasSession, err := service.CheckVersion(context.Background(), login, token, usersVersion)

	if err != nil {
		t.Errorf("CheckVersion returned an unexpected error: %v", err)
	}

	if !hasSession {
		t.Error("CheckVersion returned unexpected result: expected true, got false")
	}

	mockStorage.EXPECT().CheckVersion(login, token, usersVersion).Return(false, errors.New(""))

	hasSession, err = service.CheckVersion(context.Background(), login, token, usersVersion)

	assert.Error(t, err)
}

func TestGetVersion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMocksessionStorage(ctrl)
	mockLogger := zap.NewExample().Sugar()

	service := NewSessionService(mockStorage, mockLogger)

	login := "testuser"
	token := "token123"
	expectedVersion := uint32(3)

	mockStorage.EXPECT().GetVersion(login, token).Return(expectedVersion, nil)

	version, err := service.GetVersion(context.Background(), login, token)

	if err != nil {
		t.Errorf("GetVersion returned an unexpected error: %v", err)
	}

	if version != expectedVersion {
		t.Errorf("GetVersion returned unexpected version: expected %d, got %d", expectedVersion, version)
	}

	mockStorage.EXPECT().GetVersion(login, token).Return(expectedVersion, errors.New(""))

	version, err = service.GetVersion(context.Background(), login, token)

	assert.Error(t, err)
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

	err := service.HasSession(context.Background(), login, token)

	if err != nil {
		t.Errorf("HasSession returned an unexpected error: %v", err)
	}

	mockStorage.EXPECT().HasSession(login, token).Return(errors.New(""))

	err = service.HasSession(context.Background(), login, token)

	assert.Error(t, err)
}

func TestCheckAllUserSessionTokens(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMocksessionStorage(ctrl)
	mockLogger := zap.NewExample().Sugar()

	service := NewSessionService(mockStorage, mockLogger)

	login := "testuser"

	mockStorage.EXPECT().CheckAllUserSessionTokens(login).Return(nil)

	err := service.CheckAllUserSessionTokens(context.Background(), login)

	if err != nil {
		t.Errorf("CheckAllUserSessionTokens returned an unexpected error: %v", err)
	}

	mockStorage.EXPECT().CheckAllUserSessionTokens(login).Return(errors.New(""))

	err = service.CheckAllUserSessionTokens(context.Background(), login)

	assert.Error(t, err)
}
