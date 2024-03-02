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

func (sessionStorage *SessionStorage) Add(token string, version uint8) (err error) {
	if _, hasUser := sessionStorage.cacheStorage.Get(token); !hasUser {
		sessionStorage.cacheStorage.Set(token, 1, 0)
		return nil
	}
	return itemsIsAlreadyInTheCache
}

func (sessionStorage *SessionStorage) Delete(token string) (err error) {
	if _, hasUser := sessionStorage.cacheStorage.Get(token); hasUser {
		sessionStorage.cacheStorage.Delete(token)
		return nil
	}
	return noSuchItemInTheCache
}

func (sessionStorage *SessionStorage) Update(token string) (err error) {
	if version, hasUser := sessionStorage.cacheStorage.Get(token); hasUser {
		sessionStorage.cacheStorage.Set(token, version.(int)+1, 0)
		return nil
	}
	return noSuchItemInTheCache
}

func (sessionStorage *SessionStorage) HasUser(token string, usersVersion int) (hasSession bool, err error) {
	if version, hasUser := sessionStorage.cacheStorage.Get(token); hasUser {
		if usersVersion == version {
			return true, nil
		}
		return false, wrongSessionVersion
	}
	return false, noSuchItemInTheCache
}
