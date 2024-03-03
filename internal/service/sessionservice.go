package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	SECRET = []byte("SECRETKEY")
)

type sessionStorage interface {
	Add(login string, token string, version int) (err error)
	Delete(login string, token string) (err error)
	Update(login string, token string) (err error)
	CheckVersion(login string, token string, usersVersion int) (hasSession bool, err error)
	GetVersion(login string, token string) (version int, err error)
	HasUser(login string) bool
}

type SessionService struct {
	sessionStorage sessionStorage
}

func InitSessionService(sessionStorage sessionStorage) *SessionService {
	return &SessionService{
		sessionStorage: sessionStorage,
	}
}

func (sessionStorageService *SessionService) Add(login string, token string, version int) (err error) {
	err = sessionStorageService.sessionStorage.Add(login, token, version)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (sessionStorageService *SessionService) Delete(login string, token string) (err error) {
	err = sessionStorageService.sessionStorage.Delete(login, token)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (sessionStorageService *SessionService) Update(login string, token string) (err error) {
	err = sessionStorageService.sessionStorage.Update(login, token)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (sessionStorageService *SessionService) CheckVersion(login string, token string, usersVersion int) (hasSession bool, err error) {
	hasSession, err = sessionStorageService.sessionStorage.CheckVersion(login, token, usersVersion)
	if err != nil {
		fmt.Println(err)
		return hasSession, err
	}
	return hasSession, nil
}

func (sessionStorageService *SessionService) GetVersion(login string, token string) (version int, err error) {
	version, err = sessionStorageService.sessionStorage.GetVersion(login, token)
	if err != nil {
		fmt.Println(err)
		return version, err
	}
	return version, nil
}

func (sessionStorageService *SessionService) HasUser(login string) (hasUser bool) {
	hasUser = sessionStorageService.sessionStorage.HasUser(login)
	return hasUser
}

type customClaims struct {
	jwt.StandardClaims
	Login   string
	Status  string
	Version int
}

func (c customClaims) Valid() error {
	return c.StandardClaims.Valid()
}

func GenerateTokens(login string, status string, version int) (string, string, error) {
	accessCustomClaims := customClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			Issuer:    "NETrunnerFLIX",
		},
		Login:   login,
		Status:  status,
		Version: version,
	}

	refreshCustomClaims := customClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		},
		Login:   login,
		Status:  status,
		Version: version,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessCustomClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshCustomClaims)

	accessTokenSigned, accessErr := accessToken.SignedString(SECRET)
	refreshTokenSigned, refreshErr := refreshToken.SignedString(SECRET)
	if accessErr == nil && refreshErr == nil {
		return accessTokenSigned, refreshTokenSigned, nil
	}
	if accessErr != nil {
		return "", "", accessErr
	}
	return "", "", refreshErr
}
