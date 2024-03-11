package service

import (
	"fmt"
)

type sessionStorage interface {
	Add(login string, token string, version uint8) (err error)
	DeleteSession(login string, token string) (err error)
	Update(login string, token string) (err error)
	CheckVersion(login string, token string, usersVersion uint8) (hasSession bool, err error)
	GetVersion(login string, token string) (version uint8, err error)
	HasSession(login string, token string) error
	CheckAllUserSessionTokens(login string) error
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

func (sessionStorageService *SessionService) DeleteSession(login string, token string) (err error) {
	err = sessionStorageService.sessionStorage.DeleteSession(login, token)
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

func (sessionStorageService *SessionService) CheckVersion(login string, token string,
	usersVersion uint8) (hasSession bool, err error) {
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

func (sessionStorageService *SessionService) HasSession(login string, token string) (err error) {
	err = sessionStorageService.sessionStorage.HasSession(login, token)
	return err
}

func (sessionStorageService *SessionService) CheckAllUserSessionTokens(login string) error {
	err := sessionStorageService.sessionStorage.CheckAllUserSessionTokens(login)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
