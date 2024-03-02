package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	SECRET        = []byte("SECRETKEY")
	notPOSTMethod = errors.New("метод передачи не POST")
)

type sessionStorage interface {
	Add(token string) error
	Delete(token string) error
	Update(token string) error
	HasUser(token string, usersVersion int) (bool, error)
	GetVersion(login string, token string) (version uint8, err error)
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

func GenerateTokens(login string, status string, version int) (string, string, error) {
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
	if accessErr != nil && refreshErr != nil {
		return accessTokenSigned, refreshTokenSigned, nil
	}
	return "", "", accessErr

}
