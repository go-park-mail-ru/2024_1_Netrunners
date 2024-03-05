package cache

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/patrickmn/go-cache"

	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
)

var (
	SECRET           = os.Getenv("SECRETKEY")
	maxVersion uint8 = 255
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
		return myerrors.ErrItemsIsAlreadyInTheCache
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
			fmt.Println(sessionMapInterface.(map[string]uint8))
			return nil
		}
		return myerrors.ErrNoSuchSessionInTheCache
	}
	return myerrors.ErrNoSuchUserInTheCache
}

func (sessionStorage *SessionStorage) Update(login string, token string) (err error) {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		if _, hasSession := sessionMapInterface.(map[string]uint8)[token]; hasSession {
			if (sessionMapInterface.(map[string]uint8))[token] == maxVersion {
				return myerrors.ErrTooHighVersion
			}
			(sessionMapInterface.(map[string]uint8))[token]++
			return nil
		}
		return myerrors.ErrNoSuchSessionInTheCache
	}
	return myerrors.ErrNoSuchUserInTheCache
}

func (sessionStorage *SessionStorage) CheckVersion(login string, token string, usersVersion uint8) (hasSession bool, err error) {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		if sessionMapInterface.(map[string]uint8)[token] == usersVersion {
			return true, nil
		}
		return false, myerrors.ErrWrongSessionVersion
	}
	return false, myerrors.ErrNoSuchItemInTheCache
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
	return 0, myerrors.ErrNoSuchItemInTheCache
}

func (sessionStorage *SessionStorage) CheckAllUserSessionTokens(login string) error {
	if sessionMapInterface, hasUser := sessionStorage.cacheStorage.Get(login); hasUser {
		for token, version := range sessionMapInterface.(map[string]uint8) {
			fmt.Println("token:", token, "version:", version)
		}
		return nil
	}
	return myerrors.ErrNoSuchUserInTheCache
}

func (sessionStorage *SessionStorage) GenerateTokens(login string, status string, version uint8) (accessTokenSigned string, refreshTokenSigned string, err error) {
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
			Issuer:    "NETrunnerFLIX",
		},
		Login:   login,
		Status:  status,
		Version: version,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessCustomClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshCustomClaims)

	accessTokenSigned, err = accessToken.SignedString([]byte(SECRET))
	if err != nil {
		return "", "", fmt.Errorf("%v, %w", err, myerrors.ErrInternalServerError)
	}
	refreshTokenSigned, err = refreshToken.SignedString([]byte(SECRET))
	if err != nil {
		return "", "", fmt.Errorf("%v, %w", err, myerrors.ErrInternalServerError)
	}

	return accessTokenSigned, refreshTokenSigned, nil
}
