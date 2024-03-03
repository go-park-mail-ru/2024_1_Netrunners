package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"net/http"
	"os"
)

var (
	SECRET = os.Getenv("SECRETKEY")
)

type usersStorage interface {
	Create(user domain.User) error
	Remove(login string) error
	HasUser(login, password string) error
	GetUser(login string) (domain.User, error)
	ChangePassword(login, newPassword string) error
	ChangeName(login, newName string) (domain.User, error)
}

type AuthService struct {
	storage usersStorage
}

func InitAuthService(storage usersStorage) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (authService *AuthService) Create(user domain.User) error {
	err := authService.storage.Create(user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (authService *AuthService) Remove(login string) error {
	err := authService.storage.Remove(login)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (authService *AuthService) HasUser(login, password string) error {
	err := authService.storage.HasUser(login, password)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (authService *AuthService) GetUser(login string) (domain.User, error) {
	user, err := authService.storage.GetUser(login)
	if err != nil {
		fmt.Println(err)
		return domain.User{}, err
	}
	return user, nil
}

func (authService *AuthService) ChangePassword(login, newPassword string) error {
	err := authService.storage.ChangePassword(login, newPassword)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (authService *AuthService) ChangeName(login, newName string) (domain.User, error) {
	user, err := authService.storage.GetUser(login)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		return nil, err
	}
	return parsedToken, nil
}
