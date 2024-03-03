package cache

import (
	"errors"
	"github.com/patrickmn/go-cache"
)

var (
	itemsIsAlreadyInTheCache = errors.New("такого токена нет в кэше")
	noSuchItemInTheCache     = errors.New("такой токен уже есть в кэше")
	wrongSessionVersion      = errors.New("версии отличаются")
)

type SessionStorage struct {
	cacheStorage cache.Cache
}

func InitSessionStorage(cacheStorage *cache.Cache) *SessionStorage {
	return &SessionStorage{
		cacheStorage: *cacheStorage,
	}
}

func (sessionStorage *SessionStorage) Add(login string, token string, version int) (err error) {
	if sessionMap, hasUser := sessionStorage.cacheStorage.Get(token); !hasUser {
		sesMap := make(map[string]int)
		sesMap[token] = version
		sessionStorage.cacheStorage.Set(login, sesMap, 0)
		return nil
	} else {
		sesMap := sessionMap.(map[string]int)
		sesMap[token] = version
		sessionStorage.cacheStorage.Set(login, sesMap, 0)
	}
	return itemsIsAlreadyInTheCache
}

func (sessionStorage *SessionStorage) Delete(login string, token string) (err error) {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); !hasUser {
		delete(sessionMapInterface.(map[string]int), token)
		return nil
	}
	return noSuchItemInTheCache
}

func (sessionStorage *SessionStorage) Update(login string, token string) (err error) {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		sessionStorage.cacheStorage.Set(token, (sessionMapInterface.(map[string]int))[token]+1, 0)
		return nil
	}
	sessionMap := make(map[string]int)
	sessionMap[token] = 1
	sessionStorage.cacheStorage.Set(token, sessionMap, 0)

	return noSuchItemInTheCache
}

func (sessionStorage *SessionStorage) CheckVersion(login string, token string, usersVersion int) (hasSession bool, err error) {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		if sessionMapInterface.(map[string]int)[token] == usersVersion {
			return true, nil
		}
		return false, wrongSessionVersion
	}
	return false, noSuchItemInTheCache
}

func (sessionStorage *SessionStorage) GetVersion(login string, token string) (version int, err error) {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		return sessionMapInterface.(map[string]int)[token], nil
	}
	return 0, noSuchItemInTheCache
}
