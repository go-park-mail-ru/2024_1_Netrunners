package api_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/users/api"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/users/mocks"
)

func TestUsersServer_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersService := mocks.NewMockUsersService(ctrl)

	server := api.NewUsersServer(mockUsersService, nil)

	user := domain.UserSignUp{
		Email:    "test@example.com",
		Password: "password",
	}

	mockUsersService.EXPECT().CreateUser(gomock.Any(), user).Return(nil)

	req := &session.CreateUserRequest{User: &session.UserSignUp{
		Email:    user.Email,
		Password: user.Password,
	}}
	_, err := server.CreateUser(context.Background(), req)

	assert.NoError(t, err)
}

func TestUsersServer_RemoveUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersService := mocks.NewMockUsersService(ctrl)

	server := api.NewUsersServer(mockUsersService, nil)

	mockUsersService.EXPECT().RemoveUser(gomock.Any(), "test@example.com").Return(nil)

	req := &session.RemoveUserRequest{Login: "test@example.com"}
	_, err := server.RemoveUser(context.Background(), req)

	assert.NoError(t, err)
}

func TestUsersServer_HasUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersService := mocks.NewMockUsersService(ctrl)

	server := api.NewUsersServer(mockUsersService, nil)

	mockUsersService.EXPECT().HasUser(gomock.Any(), "test@example.com", "password").Return(nil)

	req := &session.HasUserRequest{Login: "test@example.com", Password: "password"}
	_, err := server.HasUser(context.Background(), req)

	assert.NoError(t, err)
}

func TestUsersServer_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersService := mocks.NewMockUsersService(ctrl)

	server := api.NewUsersServer(mockUsersService, nil)

	mockUser := domain.User{
		Email:    "test@example.com",
		Password: "password",
	}

	mockUsersService.EXPECT().GetUser(gomock.Any(), "test@example.com").Return(mockUser, nil)

	req := &session.GetUserRequest{Login: "test@example.com"}
	res, err := server.GetUser(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUsersServer_ChangeUserPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersService := mocks.NewMockUsersService(ctrl)

	server := api.NewUsersServer(mockUsersService, nil)

	mockUser := domain.User{
		Email:    "test@example.com",
		Password: "password",
	}

	mockUsersService.EXPECT().ChangeUserPassword(gomock.Any(), "test@example.com", "newPassword").Return(mockUser, nil)

	req := &session.ChangeUserPasswordRequest{Login: "test@example.com", NewPassword: "newPassword"}
	res, err := server.ChangeUserPassword(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUsersServer_ChangeUserName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersService := mocks.NewMockUsersService(ctrl)

	server := api.NewUsersServer(mockUsersService, nil)

	mockUser := domain.User{
		Email:    "test@example.com",
		Password: "password",
	}

	mockUsersService.EXPECT().ChangeUserName(gomock.Any(), "test@example.com", "newUsername").Return(mockUser, nil)

	req := &session.ChangeUserNameRequest{Login: "test@example.com", NewUsername: "newUsername"}
	res, err := server.ChangeUserName(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUsersServer_GetUserDataByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersService := mocks.NewMockUsersService(ctrl)

	server := api.NewUsersServer(mockUsersService, nil)

	mockUser := domain.User{
		Email:    "test@example.com",
		Password: "password",
	}

	mockUsersService.EXPECT().GetUserDataByUuid(gomock.Any(), "uuid").Return(mockUser, nil)

	req := &session.GetUserDataByUuidRequest{Uuid: "uuid"}
	res, err := server.GetUserDataByUuid(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUsersServer_GetUserPreview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersService := mocks.NewMockUsersService(ctrl)

	server := api.NewUsersServer(mockUsersService, nil)

	mockUserPreview := domain.UserPreview{
		Name: "test@example.com",
	}

	mockUsersService.EXPECT().GetUserPreview(gomock.Any(), "uuid").Return(mockUserPreview, nil)

	req := &session.GetUserPreviewRequest{Uuid: "uuid"}
	res, err := server.GetUserPreview(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUsersServer_ChangeUserPasswordByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersService := mocks.NewMockUsersService(ctrl)

	server := api.NewUsersServer(mockUsersService, nil)

	mockUser := domain.User{
		Email:    "test@example.com",
		Password: "password",
	}

	mockUsersService.EXPECT().ChangeUserPasswordByUuid(gomock.Any(), "uuid", "newPassword").Return(mockUser, nil)

	req := &session.ChangeUserPasswordByUuidRequest{Uuid: "uuid", NewPassword: "newPassword"}
	res, err := server.ChangeUserPasswordByUuid(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUsersServer_ChangeUserNameByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersService := mocks.NewMockUsersService(ctrl)

	server := api.NewUsersServer(mockUsersService, nil)

	mockUser := domain.User{
		Email:    "test@example.com",
		Password: "password",
	}

	mockUsersService.EXPECT().ChangeUserNameByUuid(gomock.Any(), "uuid", "newUsername").Return(mockUser, nil)

	req := &session.ChangeUserNameByUuidRequest{Uuid: "uuid", NewUsername: "newUsername"}
	res, err := server.ChangeUserNameByUuid(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUsersServer_ChangeUserAvatarByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsersService := mocks.NewMockUsersService(ctrl)

	server := api.NewUsersServer(mockUsersService, nil)

	mockUser := domain.User{
		Email:    "test@example.com",
		Password: "password",
	}

	mockUsersService.EXPECT().ChangeUserAvatarByUuid(gomock.Any(), "uuid", "newAvatar").Return(mockUser, nil)

	req := &session.ChangeUserAvatarByUuidRequest{Uuid: "uuid", NewAvatar: "newAvatar"}
	res, err := server.ChangeUserAvatarByUuid(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
}
