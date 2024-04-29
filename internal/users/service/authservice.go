package service

import (
	"context"
<<<<<<< HEAD
	"os"

	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
=======
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
>>>>>>> 6fcd4c8 (everything wrong)
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

<<<<<<< HEAD
type UsersService struct {
=======
type AuthService struct {
>>>>>>> 6fcd4c8 (everything wrong)
	storage   usersStorage
	secretKey string
	logger    *zap.SugaredLogger
}

<<<<<<< HEAD
func NewUsersService(storage usersStorage, logger *zap.SugaredLogger) *UsersService {
	return &UsersService{
=======
func NewAuthService(storage usersStorage, logger *zap.SugaredLogger) *AuthService {
	return &AuthService{
>>>>>>> 6fcd4c8 (everything wrong)
		storage:   storage,
		logger:    logger,
		secretKey: os.Getenv("SECRETKEY"),
	}
}

<<<<<<< HEAD
func (service *UsersService) CreateUser(ctx context.Context, user domain.UserSignUp) error {
=======
func (service *AuthService) CreateUser(ctx context.Context, user domain.UserSignUp) error {
>>>>>>> 6fcd4c8 (everything wrong)
	err := service.storage.CreateUser(user)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to create user: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return err
	}
	return nil
}

<<<<<<< HEAD
func (service *UsersService) RemoveUser(ctx context.Context, login string) error {
=======
func (service *AuthService) RemoveUser(ctx context.Context, login string) error {
>>>>>>> 6fcd4c8 (everything wrong)
	err := service.storage.RemoveUser(login)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to remove user: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return err
	}
	return nil
}

<<<<<<< HEAD
func (service *UsersService) HasUser(ctx context.Context, login, password string) error {
=======
func (service *AuthService) HasUser(ctx context.Context, login, password string) error {
>>>>>>> 6fcd4c8 (everything wrong)
	err := service.storage.HasUser(login, password)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to has user: %v", ctx.Value(requestId.ReqIDKey), err)
		return err
	}
	return nil
}

<<<<<<< HEAD
func (service *UsersService) GetUser(ctx context.Context, login string) (domain.User, error) {
=======
func (service *AuthService) GetUser(ctx context.Context, login string) (domain.User, error) {
>>>>>>> 6fcd4c8 (everything wrong)
	user, err := service.storage.GetUser(login)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get user: %v", ctx.Value(requestId.ReqIDKey), err)
		return domain.User{}, err
	}
	return user, nil
}

<<<<<<< HEAD
func (service *UsersService) ChangeUserPassword(ctx context.Context, login, newPassword string) (domain.User, error) {
=======
func (service *AuthService) ChangeUserPassword(ctx context.Context, login, newPassword string) (domain.User, error) {
>>>>>>> 6fcd4c8 (everything wrong)
	user, err := service.storage.ChangeUserPassword(login, newPassword)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to change password: %v",
			ctx.Value(requestId.ReqIDKey), err)
		return domain.User{}, err
	}
	return user, nil
}

<<<<<<< HEAD
func (service *UsersService) ChangeUserName(ctx context.Context, login, newName string) (domain.User, error) {
=======
func (service *AuthService) ChangeUserName(ctx context.Context, login, newName string) (domain.User, error) {
>>>>>>> 6fcd4c8 (everything wrong)
	user, err := service.storage.ChangeUserName(login, newName)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to change username: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return domain.User{}, err
	}
	return user, nil
}

<<<<<<< HEAD
func (service *UsersService) GetUserDataByUuid(ctx context.Context, uuid string) (domain.User, error) {
=======
func (service *AuthService) IsTokenValid(token *http.Cookie) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(service.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token: %w", myerrors.ErrNotAuthorised)
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token: %w", myerrors.ErrNotAuthorised)
	}

	_, ok = claims["Login"]
	if !ok {
		return nil, fmt.Errorf("invalid token: %w", myerrors.ErrNotAuthorised)
	}
	_, ok = claims["IsAdmin"]
	if !ok {
		return nil, fmt.Errorf("invalid token: %w", myerrors.ErrNotAuthorised)
	}

	_, ok = claims["Version"]
	if !ok {
		return nil, fmt.Errorf("invalid token: %w", myerrors.ErrNotAuthorised)
	}

	return claims, nil
}

func ValidateLogin(e string) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if emailRegex.MatchString(e) {
		return nil
	}
	return myerrors.ErrLoginIsNotValid
}

func ValidateUsername(username string) error {
	if len(username) >= 4 {
		return nil
	}
	return myerrors.ErrUsernameIsToShort
}

func ValidatePassword(password string) error {
	if len(password) >= 6 {
		return nil
	}
	return myerrors.ErrPasswordIsToShort
}

type customClaims struct {
	jwt.StandardClaims
	Login   string
	IsAdmin bool
	Version uint8
}

func (service *AuthService) GenerateTokens(login string, isAdmin bool, version uint8) (tokenSigned string, err error) {
	tokenCustomClaims := customClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    "NETrunnerFLIX",
		},
		Login:   login,
		IsAdmin: isAdmin,
		Version: version,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenCustomClaims)

	tokenSigned, err = token.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	return tokenSigned, nil
}

func (service *AuthService) GetUserDataByUuid(ctx context.Context, uuid string) (domain.User, error) {
>>>>>>> 6fcd4c8 (everything wrong)
	user, err := service.storage.GetUserDataByUuid(uuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get user data: %v",
			ctx.Value(requestId.ReqIDKey), err)
		return domain.User{}, err
	}
	return user, nil
}

<<<<<<< HEAD
func (service *UsersService) GetUserPreview(ctx context.Context, uuid string) (domain.UserPreview, error) {
=======
func (service *AuthService) GetUserPreview(ctx context.Context, uuid string) (domain.UserPreview, error) {
>>>>>>> 6fcd4c8 (everything wrong)
	userPreview, err := service.storage.GetUserPreview(uuid)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get user preview: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return domain.UserPreview{}, err
	}
	return userPreview, nil
}

<<<<<<< HEAD
func (service *UsersService) ChangeUserPasswordByUuid(ctx context.Context, uuid, newPassword string) (domain.User,
=======
func (service *AuthService) ChangeUserPasswordByUuid(ctx context.Context, uuid, newPassword string) (domain.User,
>>>>>>> 6fcd4c8 (everything wrong)
	error) {
	user, err := service.storage.ChangeUserPasswordByUuid(uuid, newPassword)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to change password: %v",
			ctx.Value(requestId.ReqIDKey), err)
		return domain.User{}, err
	}
	return user, nil
}

<<<<<<< HEAD
func (service *UsersService) ChangeUserNameByUuid(ctx context.Context, uuid, newName string) (domain.User, error) {
=======
func (service *AuthService) ChangeUserNameByUuid(ctx context.Context, uuid, newName string) (domain.User, error) {
>>>>>>> 6fcd4c8 (everything wrong)
	user, err := service.storage.ChangeUserNameByUuid(uuid, newName)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to change username: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return domain.User{}, err
	}
	return user, nil
}
<<<<<<< HEAD

func (service *UsersService) ChangeUserAvatarByUuid(ctx context.Context, uuid, newAvatar string) (domain.User, error) {
	user, err := service.storage.ChangeUserAvatarByUuid(uuid, newAvatar)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to change username: %v", ctx.Value(requestId.ReqIDKey),
			err)
		return domain.User{}, err
	}
	return user, nil
}
=======
>>>>>>> 6fcd4c8 (everything wrong)
