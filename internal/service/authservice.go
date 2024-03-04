package service

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
)

var (
	SECRET = os.Getenv("SECRETKEY")
)

type usersStorage interface {
	CreateUser(user domain.User) error
	RemoveUser(login string) error
	HasUser(login, password string) error
	GetUser(login string) (domain.User, error)
	ChangeUserPassword(login, newPassword string) error
	ChangeUserName(login, newName string) (domain.User, error)
}

type AuthService struct {
	storage usersStorage
}

func InitAuthService(storage usersStorage) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (authService *AuthService) CreateUser(user domain.User) error {
	err := authService.storage.CreateUser(user)
	if err != nil {
		fmt.Printf("creating user error: %v", err)
		return err
	}
	return nil
}

func (authService *AuthService) RemoveUser(login string) error {
	err := authService.storage.RemoveUser(login)
	if err != nil {
		fmt.Printf("creating user error: %v", err)
		return err
	}
	return nil
}

func (authService *AuthService) HasUser(login, password string) error {
	err := authService.storage.HasUser(login, password)
	if err != nil {
		fmt.Printf("creating user error: %v", err)
		return err
	}
	return nil
}

func (authService *AuthService) GetUser(login string) (domain.User, error) {
	user, err := authService.storage.GetUser(login)
	if err != nil {
		fmt.Printf("creating user error: %v", err)
		return domain.User{}, err
	}
	return user, nil
}

func (authService *AuthService) ChangeUserPassword(login, newPassword string) error {
	err := authService.storage.ChangeUserPassword(login, newPassword)
	if err != nil {
		fmt.Printf("creating user error: %v", err)
		return err
	}
	return nil
}

func (authService *AuthService) ChangeUserName(login, newName string) (domain.User, error) {
	user, err := authService.storage.GetUser(login)
	if err != nil {
		fmt.Printf("creating user error: %v", err)
		return domain.User{}, err
	}
	return user, nil
}

func (authService *AuthService) IsTokenValid(token *http.Cookie) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SECRET), nil
	})
	if err != nil {
		fmt.Printf("creating user error: %v", err)
		return nil, err
	}
	return parsedToken, nil
}
