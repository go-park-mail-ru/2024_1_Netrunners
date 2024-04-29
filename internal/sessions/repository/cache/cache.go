package cache

import (
	"fmt"

	"github.com/patrickmn/go-cache"

	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
)

var (
	maxVersion uint32 = 255
)

type SessionStorage struct {
	cacheStorage *cache.Cache
}

func NewSessionStorage() *SessionStorage {
	return &SessionStorage{
		cacheStorage: cache.New(0, 0),
	}
}

func (sessionStorage *SessionStorage) Add(login string, token string, version uint32) (err error) {
	sessionMap, hasUser := sessionStorage.cacheStorage.Get(login)
	if !hasUser {
		sesMap := make(map[string]uint32)
		sesMap[token] = version
		sessionStorage.cacheStorage.Set(login, sesMap, 0)
		return nil
	}
	if _, hasSession := sessionMap.(map[string]uint32)[token]; hasSession {
		return myerrors.ErrItemsIsAlreadyInTheCache

	}
	sesMap := sessionMap.(map[string]uint32)
	sesMap[token] = version
	sessionStorage.cacheStorage.Set(login, sesMap, 0)

	return nil
}

func (sessionStorage *SessionStorage) DeleteSession(login string, token string) (err error) {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		if _, hasSession := sessionMapInterface.(map[string]uint32)[token]; hasSession {
			delete(sessionMapInterface.(map[string]uint32), token)
			return nil
		}
		return myerrors.ErrNoSuchSessionInTheCache
	}
	return myerrors.ErrNoSuchUserInTheCache
}

func (sessionStorage *SessionStorage) Update(login string, token string) (err error) {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		if _, hasSession := sessionMapInterface.(map[string]uint32)[token]; hasSession {
			if (sessionMapInterface.(map[string]uint32))[token] == maxVersion {
				return myerrors.ErrTooHighVersion
			}
			(sessionMapInterface.(map[string]uint32))[token]++
			return nil
		}
		return myerrors.ErrNoSuchSessionInTheCache
	}
	return myerrors.ErrNoSuchUserInTheCache
}

func (sessionStorage *SessionStorage) CheckVersion(login string, token string,
	usersVersion uint32) (hasSession bool, err error) {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		if sessionMapInterface.(map[string]uint32)[token] == usersVersion {
			return true, nil
		}
		return false, myerrors.ErrWrongSessionVersion
	}
	return false, myerrors.ErrNoSuchItemInTheCache
}

func (sessionStorage *SessionStorage) HasSession(login string, token string) error {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		if _, hasSession := sessionMapInterface.(map[string]uint32)[token]; hasSession {
			return nil
		}
	}
	return myerrors.ErrNoSuchUser
}

func (sessionStorage *SessionStorage) GetVersion(login string, token string) (version uint32, err error) {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		return sessionMapInterface.(map[string]uint32)[token], nil
	}
	return 0, myerrors.ErrNoSuchItemInTheCache
}

func (sessionStorage *SessionStorage) CheckAllUserSessionTokens(login string) error {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		for token, version := range sessionMapInterface.(map[string]uint32) {
			fmt.Println("token:", token, "version:", version)
		}
		return nil
	}
	return myerrors.ErrNoSuchUserInTheCache
}
