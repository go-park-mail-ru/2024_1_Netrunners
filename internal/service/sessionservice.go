package service

import (
	"fmt"
)

type sessionStorage interface {
	Add(login string, token string, version uint8) (err error)
	Delete(login string, token string) (err error)
	Update(login string, token string) (err error)
	CheckVersion(login string, token string, usersVersion uint8) (hasSession bool, err error)
	GetVersion(login string, token string) (version uint8, err error)
	HasUser(login string) bool
	GenerateTokens(login string, status string, version uint8) (string, string, error)
}

type SessionService struct {
	sessionStorage sessionStorage
}

func InitSessionService(sessionStorage sessionStorage) *SessionService {
	return &SessionService{
		sessionStorage: sessionStorage,
	}
}

func (sessionStorageService *SessionService) Add(login string, token string, version uint8) (err error) {
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

func (sessionStorageService *SessionService) CheckVersion(login string, token string, usersVersion uint8) (hasSession bool, err error) {
	hasSession, err = sessionStorageService.sessionStorage.CheckVersion(login, token, usersVersion)
	if err != nil {
		fmt.Println(err)
		return hasSession, err
	}
	return hasSession, nil
}

func (sessionStorageService *SessionService) GetVersion(login string, token string) (version uint8, err error) {
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

func (sessionStorageService *SessionService) GenerateTokens(login string, status string, version uint8) (string, string, error) {
	accessToken, refreshToken, err := sessionStorageService.sessionStorage.GenerateTokens(login, status, version)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
