package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// у юзера в куке лежит только два токена аксесс (5 минут) и рефреш ( 48 часов ). мы делаем конкретно ручку для
// проверки валидности сессии. Рефреш передается только на ручки /auth.

// если куки нет, то генерим новую на основе логина и статуса юзера.
// если кука есть, проверяем на валидность, если не валидная, проверяем рефрешТокен на валидность, если валиден, то
// 		создаем новый рефреш и аксес, и отправляем их
// если кука есть, проверяем на валидность, если не валидная, проверяем рефрешТокен на валидность, если невалиден,
//		то возвращаем 401
//

var (
	SECRET        = []byte("SECRETKEY")
	notPOSTMethod = errors.New("метод передачи не POST")
)

type sessionStorage interface {
	Add(token string) error
	Delete(token string) error
	Update(token string) error
	HasUser(token string, usersVersion int) (bool, error)
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

func GenerateTokens(existingTokenString string, login string, status string) (string, string, error) {
	token, err := jwt.Parse(existingTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return SECRET, nil
	})

	version := 1

	if err == nil {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return "", "", fmt.Errorf("failed to parse claims")
		}

		login, ok = claims["login"].(string)
		if !ok {
			return "", "", fmt.Errorf("failed to parse username")
		}
		status, ok = claims["status"].(string)
		if !ok {
			return "", "", fmt.Errorf("failed to parse status")
		}
		version, ok = claims["version"].(int)
		if !ok {
			return "", "", fmt.Errorf("failed to parse version")
		}
	}

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
		Login: login,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessCustomClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshCustomClaims)

	accessTokenSigned, accessErr := accessToken.SignedString(SECRET)
	refreshTokenSigned, refreshErr := refreshToken.SignedString(SECRET)
	if accessErr != nil && refreshErr != nil {
		return accessTokenSigned, refreshTokenSigned, nil
	}
	return "", "", accessErr

}
