package cache

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/patrickmn/go-cache"
	"os"
	"time"
)

var (
	noSuchItemInTheCache           = errors.New("no such token in cache")
	itemsIsAlreadyInTheCache       = errors.New("such token is already in the cache")
	wrongSessionVersion            = errors.New("different versions")
	noSuchSessionInTheCache        = errors.New("no such session in cache")
	noSuchUserInTheCache           = errors.New("no such user in cache")
	tooHighVersion                 = errors.New("too high session version")
	SECRET                         = os.Getenv("SECRETKEY")
	maxVersion               uint8 = 255
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
	sessionMap, hasUser := sessionStorage.cacheStorage.Get(login)
	if !hasUser {
		sesMap := make(map[string]uint8)
		sesMap[token] = version
		sessionStorage.cacheStorage.Set(login, sesMap, 0)
		return nil
	}
	if _, hasSession := sessionMap.(map[string]uint8)[token]; hasSession {
		return itemsIsAlreadyInTheCache
	}
	sesMap := sessionMap.(map[string]uint8)
	sesMap[token] = version
	sessionStorage.cacheStorage.Set(login, sesMap, 0)

	return nil
}

func (sessionStorage *SessionStorage) DeleteSession(login string, token string) (err error) {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		if _, hasSession := sessionMapInterface.(map[string]uint8)[token]; hasSession {
			delete(sessionMapInterface.(map[string]uint8), token)
			return nil
		}
		return noSuchSessionInTheCache
	}
	return noSuchUserInTheCache

}

func (sessionStorage *SessionStorage) Update(login string, token string) (err error) {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		if _, hasSession := sessionMapInterface.(map[string]uint8)[token]; hasSession {
			if (sessionMapInterface.(map[string]uint8))[token] == maxVersion {
				return tooHighVersion
			}
			(sessionMapInterface.(map[string]uint8))[token]++
			return nil
		}
		return noSuchSessionInTheCache
	}
	return noSuchUserInTheCache
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

func (sessionStorage *SessionStorage) HasSession(login string, token string) bool {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		if _, hasSession := sessionMapInterface.(map[string]uint8)[token]; hasSession {
			return true
		}
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
	fmt.Println(accessToken, refreshToken)

	fmt.Println(SECRET)
	accessTokenSigned, accessErr := accessToken.SignedString([]byte(SECRET))
	if accessErr != nil {
		return "", "", accessErr
	}
	refreshTokenSigned, refreshErr := refreshToken.SignedString([]byte(SECRET))
	if refreshErr != nil {
		return "", "", refreshErr
	}

	return accessTokenSigned, refreshTokenSigned, nil
}
