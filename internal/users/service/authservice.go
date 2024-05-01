package service

import (
	"context"
	"go.uber.org/zap"
	"os"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
)

type usersStorage interface {
	CreateUser(user domain.UserSignUp) error
	RemoveUser(email string) error
	HasUser(email, password string) error
	GetUser(email string) (domain.User, error)
	ChangeUserPassword(email, newPassword string) (domain.User, error)
	ChangeUserName(email, newName string) (domain.User, error)
	GetUserDataByUuid(uuid string) (domain.User, error)
	GetUserPreview(uuid string) (domain.UserPreview, error)
	ChangeUserPasswordByUuid(uuid, newPassword string) (domain.User, error)
	ChangeUserNameByUuid(uuid, newName string) (domain.User, error)
	ChangeUserAvatarByUuid(uuid, filename string) (domain.User, error)
}

type UsersService struct {
	storage   usersStorage
	secretKey string
	logger    *zap.SugaredLogger
}

func NewUsersService(storage usersStorage, logger *zap.SugaredLogger) *UsersService {
	return &UsersService{
		storage:   storage,
		logger:    logger,
		secretKey: os.Getenv("SECRETKEY"),
	}
}

func (service *UsersService) CreateUser(ctx context.Context, user domain.UserSignUp) error {
	err := service.storage.CreateUser(user)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to create user: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return err
	}
	return nil
}

func (service *UsersService) RemoveUser(ctx context.Context, login string) error {
	err := service.storage.RemoveUser(login)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to remove user: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return err
	}
	return nil
}

func (service *UsersService) HasUser(ctx context.Context, login, password string) error {
	err := service.storage.HasUser(login, password)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to has user: %v", ctx.Value(requestId.ReqIDKey), err)
		return err
	}
	return nil
}

func (service *UsersService) GetUser(ctx context.Context, login string) (domain.User, error) {
	user, err := service.storage.GetUser(login)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get user: %v", ctx.Value(requestId.ReqIDKey), err)
		return domain.User{}, err
	}
	return user, nil
}

func (service *UsersService) ChangeUserPassword(ctx context.Context, login, newPassword string) (domain.User, error) {
	user, err := service.storage.ChangeUserPassword(login, newPassword)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to change password: %v",
			ctx.Value(requestId.ReqIDKey), err)
		return domain.User{}, err
	}
	return user, nil
}

func (service *UsersService) ChangeUserName(ctx context.Context, login, newName string) (domain.User, error) {
	user, err := service.storage.ChangeUserName(login, newName)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to change username: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return domain.User{}, err
	}
	return user, nil
}

func (service *UsersService) GetUserDataByUuid(ctx context.Context, uuid string) (domain.User, error) {
	user, err := service.storage.GetUserDataByUuid(uuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get user data: %v",
			ctx.Value(requestId.ReqIDKey), err)
		return domain.User{}, err
	}
	return user, nil
}

func (service *UsersService) GetUserPreview(ctx context.Context, uuid string) (domain.UserPreview, error) {
	userPreview, err := service.storage.GetUserPreview(uuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get user preview: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return domain.UserPreview{}, err
	}
	return userPreview, nil
}

func (service *UsersService) ChangeUserPasswordByUuid(ctx context.Context, uuid, newPassword string) (domain.User,
	error) {
	user, err := service.storage.ChangeUserPasswordByUuid(uuid, newPassword)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to change password: %v",
			ctx.Value(requestId.ReqIDKey), err)
		return domain.User{}, err
	}
	return user, nil
}

func (service *UsersService) ChangeUserNameByUuid(ctx context.Context, uuid, newName string) (domain.User, error) {
	user, err := service.storage.ChangeUserNameByUuid(uuid, newName)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to change username: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return domain.User{}, err
	}
	return user, nil
}

func (service *UsersService) ChangeUserAvatarByUuid(ctx context.Context, uuid, newAvatar string) (domain.User, error) {
	user, err := service.storage.ChangeUserAvatarByUuid(uuid, newAvatar)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to change username: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return domain.User{}, err
	}
	return user, nil
}
