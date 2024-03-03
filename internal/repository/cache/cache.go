package cache

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/patrickmn/go-cache"
	"os"
	"time"
)

var (
	noSuchItemInTheCache     = errors.New("no such token in cache")
	itemsIsAlreadyInTheCache = errors.New("such token is already in the cache")
	wrongSessionVersion      = errors.New("different versions")
	SECRET                   = os.Getenv("SECRETKEY")
)

type SessionStorage struct {
	cacheStorage *cache.Cache
}

type customClaims struct {
	jwt.StandardClaims
	Login   string
	Status  string
	Version uint8
}

func InitSessionStorage() *SessionStorage {
	return &SessionStorage{
		cacheStorage: cache.New(0, 0),
	}
}

func (sessionStorage *SessionStorage) Add(login string, token string, version uint8) (err error) {
	if sessionMap, hasUser := sessionStorage.cacheStorage.Get(token); !hasUser {
		sesMap := make(map[string]uint8)
		sesMap[token] = version
		sessionStorage.cacheStorage.Set(login, sesMap, 0)
		return nil
	} else {
		sesMap := sessionMap.(map[string]uint8)
		sesMap[token] = version
		sessionStorage.cacheStorage.Set(login, sesMap, 0)
	}
	return itemsIsAlreadyInTheCache
}

func (sessionStorage *SessionStorage) Delete(login string, token string) (err error) {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		delete(sessionMapInterface.(map[string]uint8), token)
		return nil
	}
	return noSuchItemInTheCache
}

func (sessionStorage *SessionStorage) Update(login string, token string) (err error) {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		sessionStorage.cacheStorage.Set(token, (sessionMapInterface.(map[string]uint8))[token]+1, 0)
		return nil
	}

	return noSuchItemInTheCache
}

func (sessionStorage *SessionStorage) CheckVersion(login string, token string, usersVersion uint8) (hasSession bool, err error) {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		if sessionMapInterface.(map[string]uint8)[token] == usersVersion {
			return true, nil
		}
		return false, wrongSessionVersion
	}
	return false, noSuchItemInTheCache
}

func (sessionStorage *SessionStorage) HasUser(login string) bool {
	if _, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		return true
	}
	return false
}

func (sessionStorage *SessionStorage) GetVersion(login string, token string) (version uint8, err error) {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		return sessionMapInterface.(map[string]uint8)[token], nil
	}
	return 0, noSuchItemInTheCache
}

func (sessionStorage *SessionStorage) GenerateTokens(login string, status string, version uint8) (string, string, error) {
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

	if accessErr != nil {
		return "", "", accessErr
	}
	if refreshErr != nil {
		return "", "", refreshErr
	}

	return accessTokenSigned, refreshTokenSigned, nil
}
