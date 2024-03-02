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
	cacheStorage map[string]cache.Cache
}

func InitSessionStorage(cacheStorage *map[string]cache.Cache) *SessionStorage {
	return &SessionStorage{
		cacheStorage: *cacheStorage,
	}
}

func (sessionStorage *SessionStorage) Add(login string, token string, version int) (err error) {
	if _, hasUser := sessionStorage.cacheStorage[login].Get(token); !hasUser {
		sessionStorage.cacheStorage[login].Set(token, version, 0)
		return nil
	}
	return itemsIsAlreadyInTheCache
}

func (sessionStorage *SessionStorage) Delete(login string, token string) (err error) {
	if _, hasUser := sessionStorage.cacheStorage[login].Get(token); hasUser {
		sessionStorage.cacheStorage[login].Delete(token)
		return nil
	}
	return noSuchItemInTheCache
}

func (sessionStorage *SessionStorage) Update(login string, token string) (err error) {
	if version, hasUser := sessionStorage.cacheStorage[login].Get(token); hasUser {
		sessionStorage.cacheStorage[login].Set(token, version.(int)+1, 0)
		return nil
	}
	return noSuchItemInTheCache
}

func (sessionStorage *SessionStorage) HasUser(login string, token string, usersVersion int) (hasSession bool, err error) {
	if version, hasUser := sessionStorage.cacheStorage[login].Get(token); hasUser {
		if usersVersion == version {
			return true, nil
		}
		return false, wrongSessionVersion
	}
	return false, noSuchItemInTheCache
}

func (sessionStorage *SessionStorage) GetVersion(login string, token string) (version int, err error) {
	if version, hasUser := sessionStorage.cacheStorage[login].Get(token); hasUser {
		return version.(int), nil
	}
	return 0, noSuchItemInTheCache
}
