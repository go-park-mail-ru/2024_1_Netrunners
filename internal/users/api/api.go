package api

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
)

type UsersService interface {
	CreateUser(ctx context.Context, user domain.UserSignUp) error
	RemoveUser(ctx context.Context, email string) error
	HasUser(ctx context.Context, email, password string) error
	GetUser(ctx context.Context, email string) (domain.User, error)
	ChangeUserPassword(ctx context.Context, email, newPassword string) (domain.User, error)
	ChangeUserName(ctx context.Context, email, newName string) (domain.User, error)
	GetUserDataByUuid(ctx context.Context, uuid string) (domain.User, error)
	GetUserPreview(ctx context.Context, uuid string) (domain.UserPreview, error)
	ChangeUserPasswordByUuid(ctx context.Context, uuid, newPassword string) (domain.User, error)
	ChangeUserNameByUuid(ctx context.Context, uuid, newName string) (domain.User, error)
	IsTokenValid(token *http.Cookie) (jwt.MapClaims, error)
	GenerateTokens(login string, isAdmin bool, version uint8) (tokenSigned string, err error)
}

type UsersServer struct {
	usersService UsersService
	logger       *zap.SugaredLogger
}

func NewUsersServer(service UsersService, logger *zap.SugaredLogger) *UsersServer {
	return &UsersServer{
		usersService: service,
		logger:       logger,
	}
}

func (server *UsersServer) CreateUser(ctx context.Context, user domain.UserSignUp) error {
}
