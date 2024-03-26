package service

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
)

type usersStorage interface {
	CreateUser(user domain.UserSignUp) error
	RemoveUser(email string) error
	HasUser(email, password string) error
	GetUser(email string) (domain.User, error)
	ChangeUserPassword(email, newPassword string) error
	ChangeUserName(email, newName string) (domain.User, error)
	GetUserDataByUuid(uuid string) (domain.User, error)
	GetUserPreview(uuid string) (domain.UserPreview, error)
}

type AuthService struct {
	storage   usersStorage
	secretKey string
	logger    *zap.SugaredLogger
}

func NewAuthService(storage usersStorage, logger *zap.SugaredLogger) *AuthService {
	return &AuthService{
		storage:   storage,
		logger:    logger,
		secretKey: os.Getenv("SECRETKEY"),
	}
}

func (service *AuthService) CreateUser(user domain.UserSignUp) error {
	err := service.storage.CreateUser(user)
	if err != nil {
		service.logger.Errorf("service error at CreateUser: %v", err)
		return err
	}
	return nil
}

func (service *AuthService) RemoveUser(login string) error {
	err := service.storage.RemoveUser(login)
	if err != nil {
		service.logger.Errorf("service error at RemoveUser: %v", myerrors.ErrInternalServerError)
		return err
	}
	return nil
}

func (service *AuthService) HasUser(login, password string) error {
	err := service.storage.HasUser(login, password)
	if err != nil {
		service.logger.Errorf("service error at HasUser: %v", myerrors.ErrInternalServerError)
		return err
	}
	return nil
}

func (service *AuthService) GetUser(login string) (domain.User, error) {
	user, err := service.storage.GetUser(login)
	if err != nil {
		service.logger.Errorf("service error at GetUser: %v", myerrors.ErrInternalServerError)
		return domain.User{}, err
	}
	return user, nil
}

func (service *AuthService) ChangeUserPassword(login, newPassword string) error {
	err := service.storage.ChangeUserPassword(login, newPassword)
	if err != nil {
		service.logger.Errorf("service error at ChangeUserPassword: %v", myerrors.ErrInternalServerError)
		return err
	}
	return nil
}

func (service *AuthService) ChangeUserName(login, newName string) (domain.User, error) {
	user, err := service.storage.ChangeUserName(login, newName)
	if err != nil {
		service.logger.Errorf("service error at ChangeUserName: %v", myerrors.ErrInternalServerError)
		return domain.User{}, err
	}
	return user, nil
}

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
	fmt.Println(e)
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

func (service *AuthService) GenerateTokens(login string, isAdmin bool, version uint8) (accessTokenSigned string,
	err error) {
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

	tokenSigned, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", fmt.Errorf("%v, %w", err, myerrors.ErrInternalServerError)
	}

	return tokenSigned, nil
}

func (service *AuthService) GetUserDataByUuid(uuid string) (domain.User, error) {
	user, err := service.storage.GetUserDataByUuid(uuid)
	if err != nil {
		service.logger.Errorf("service error at GetUserDataByUuid: %v", myerrors.ErrInternalServerError)
		return domain.User{}, err
	}
	return user, nil
}

func (service *AuthService) GetUserPreview(uuid string) (domain.UserPreview, error) {
	userPreview, err := service.storage.GetUserPreview(uuid)
	if err != nil {
		service.logger.Errorf("service error at GetUserPreview: %v", myerrors.ErrInternalServerError)
		return domain.UserPreview{}, err
	}
	return userPreview, nil
}
