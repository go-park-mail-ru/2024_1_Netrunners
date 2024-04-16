package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
	mockService "github.com/go-park-mail-ru/2024_1_Netrunners/internal/service/mock"
)

func TestAuthService_HasUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMockusersStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	login := "cakethefake@gmail.com"
	password := "123456789"

	mockStorage.EXPECT().HasUser(login, password).Return(nil)

	authService := NewAuthService(mockStorage, mockLogger)
	err := authService.HasUser(context.Background(), login, password)

	assert.NoError(t, err)

	mockStorage.EXPECT().HasUser(login, password).Return(errors.New(""))

	authService = NewAuthService(mockStorage, mockLogger)
	err = authService.HasUser(context.Background(), login, password)

	assert.Error(t, err)
}

func TestAuthService_ChangeUserPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMockusersStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	login := "cakethefake@gmail.com"
	newPassword := "newPassword123"

	mockStorage.EXPECT().ChangeUserPassword(login, newPassword).Return(domain.User{}, nil)

	authService := NewAuthService(mockStorage, mockLogger)
	_, err := authService.ChangeUserPassword(context.Background(), login, newPassword)

	assert.NoError(t, err)

	mockStorage.EXPECT().ChangeUserPassword(login, newPassword).Return(domain.User{}, errors.New(""))

	authService = NewAuthService(mockStorage, mockLogger)
	_, err = authService.ChangeUserPassword(context.Background(), login, newPassword)

	assert.Error(t, err)
}

func TestAuthService_ChangeUserName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMockusersStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	login := "cakethefake@gmail.com"
	newName := "New Name"

	mockStorage.EXPECT().ChangeUserName(login, newName).Return(domain.User{}, nil)

	authService := NewAuthService(mockStorage, mockLogger)
	_, err := authService.ChangeUserName(context.Background(), login, newName)

	assert.NoError(t, err)
}

func TestAuthService_ValidateLogin_Valid(t *testing.T) {
	email := "cakethefake@gmail.com"
	err := ValidateLogin(email)
	assert.NoError(t, err)
}

func TestAuthService_ValidateLogin_Invalid(t *testing.T) {
	email := "invalidemail.com"
	err := ValidateLogin(email)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, myerrors.ErrLoginIsNotValid))
}

func TestAuthService_ValidateUsername_Valid(t *testing.T) {
	username := "TestUsername"
	err := ValidateUsername(username)
	assert.NoError(t, err)
}

func TestAuthService_ValidateUsername_Invalid(t *testing.T) {
	username := "Usr"
	err := ValidateUsername(username)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, myerrors.ErrUsernameIsToShort))
}

func TestAuthService_ValidatePassword_Valid(t *testing.T) {
	password := "StrongPassword123"
	err := ValidatePassword(password)
	assert.NoError(t, err)
}

func TestAuthService_ValidatePassword_Invalid(t *testing.T) {
	password := "weak"
	err := ValidatePassword(password)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, myerrors.ErrPasswordIsToShort))
}

func TestAuthService_GenerateTokens(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := zaptest.NewLogger(t).Sugar()

	authService := NewAuthService(nil, mockLogger)

	login := "cakethefake@gmail.com"
	isAdmin := true
	version := uint8(1)

	tokenSigned, err := authService.GenerateTokens(login, isAdmin, version)

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenSigned)
}

func TestAuthService_GetUserDataByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMockusersStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	uuid := "1"

	user := domain.User{
		Uuid:         uuid,
		Email:        "cakethefake@gmail.com",
		Avatar:       "",
		Name:         "Danya",
		Password:     "123456789",
		IsAdmin:      true,
		RegisteredAt: time.Now(),
		Birthday:     time.Now(),
	}

	mockStorage.EXPECT().GetUserDataByUuid(uuid).Return(user, nil)

	authService := NewAuthService(mockStorage, mockLogger)
	retrievedUser, err := authService.GetUserDataByUuid(context.Background(), uuid)

	assert.NoError(t, err)
	assert.Equal(t, user, retrievedUser)

	mockStorage.EXPECT().GetUserDataByUuid(uuid).Return(user, errors.New(""))

	authService = NewAuthService(mockStorage, mockLogger)
	_, err = authService.GetUserDataByUuid(context.Background(), uuid)

	assert.Error(t, err)
}

func TestAuthService_GetUserPreview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMockusersStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	uuid := "1"

	userPreview := domain.UserPreview{
		Name:   "Danya",
		Avatar: "",
	}

	mockStorage.EXPECT().GetUserPreview(uuid).Return(userPreview, nil)

	authService := NewAuthService(mockStorage, mockLogger)
	retrievedUserPreview, err := authService.GetUserPreview(context.Background(), uuid)

	assert.NoError(t, err)
	assert.Equal(t, userPreview, retrievedUserPreview)

	mockStorage.EXPECT().GetUserPreview(uuid).Return(userPreview, errors.New(""))

	authService = NewAuthService(mockStorage, mockLogger)
	retrievedUserPreview, err = authService.GetUserPreview(context.Background(), uuid)

	assert.Error(t, err)
}

func TestAuthService_ChangeUserPasswordByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMockusersStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	uuid := "1"
	newPassword := "newPassword123"

	mockStorage.EXPECT().ChangeUserPasswordByUuid(uuid, newPassword).Return(domain.User{}, nil)

	authService := NewAuthService(mockStorage, mockLogger)
	_, err := authService.ChangeUserPasswordByUuid(context.Background(), uuid, newPassword)

	assert.NoError(t, err)
}

func TestAuthService_ChangeUserNameByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMockusersStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	uuid := "1"
	newName := "New Name"

	mockStorage.EXPECT().ChangeUserNameByUuid(uuid, newName).Return(domain.User{}, nil)

	authService := NewAuthService(mockStorage, mockLogger)
	_, err := authService.ChangeUserNameByUuid(context.Background(), uuid, newName)

	assert.NoError(t, err)
}

func TestAuthService_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMockusersStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	login := "cakethefake@gmail.com"
	expectedUser := domain.User{Name: "Test User"}

	mockStorage.EXPECT().GetUser(gomock.Eq(login)).Return(expectedUser, nil)

	authService := NewAuthService(mockStorage, mockLogger)
	user, err := authService.GetUser(context.Background(), login)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestAuthService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMockusersStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	user := domain.UserSignUp{
		Email:    "test@example.com",
		Name:     "Test User",
		Password: "password",
	}

	mockStorage.EXPECT().CreateUser(user).Return(nil)

	authService := NewAuthService(mockStorage, mockLogger)
	err := authService.CreateUser(context.Background(), user)

	assert.NoError(t, err)
}

func TestAuthService_RemoveUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mockService.NewMockusersStorage(ctrl)
	mockLogger := zaptest.NewLogger(t).Sugar()

	login := "test@example.com"

	mockStorage.EXPECT().RemoveUser(login).Return(nil)

	authService := NewAuthService(mockStorage, mockLogger)
	err := authService.RemoveUser(context.Background(), login)

	assert.NoError(t, err)
}
