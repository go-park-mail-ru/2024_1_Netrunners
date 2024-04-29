package service

import (
	"context"
	"fmt"
<<<<<<< HEAD
<<<<<<< HEAD:internal/sessions/service/sessionservice.go
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"

=======
	"net/http"
=======
>>>>>>> d726255 (all microservices r done)
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"

<<<<<<< HEAD
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
>>>>>>> a871897 (users done, waits for sessions):internal/service/sessionservice.go
=======
>>>>>>> d726255 (all microservices r done)
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
)

type sessionStorage interface {
	Add(login string, token string, version uint32) (err error)
	DeleteSession(login string, token string) (err error)
	Update(login string, token string) (err error)
	CheckVersion(login string, token string, usersVersion uint32) (hasSession bool, err error)
	GetVersion(login string, token string) (version uint32, err error)
	HasSession(login string, token string) error
	CheckAllUserSessionTokens(login string) error
}

type SessionService struct {
	sessionStorage sessionStorage
	logger         *zap.SugaredLogger
	secretKey      string
}

func NewSessionService(sessionStorage sessionStorage, logger *zap.SugaredLogger) *SessionService {
	return &SessionService{
		sessionStorage: sessionStorage,
		logger:         logger,
		secretKey:      os.Getenv("SECRETKEY"),
	}
}

func (service *SessionService) Add(ctx context.Context, login, token string, version uint32) (err error) {
	err = service.sessionStorage.Add(login, token, version)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to add session: %v", ctx.Value(requestId.ReqIDKey), err)
		return err
	}
	return nil
}

func (service *SessionService) DeleteSession(ctx context.Context, login, token string) (err error) {
	err = service.sessionStorage.DeleteSession(login, token)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to delete session: %v",
			ctx.Value(requestId.ReqIDKey), err)
		return err
	}
	return nil
}

func (service *SessionService) Update(ctx context.Context, login, token string) (err error) {
	err = service.sessionStorage.Update(login, token)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to update session: %v", ctx.Value(requestId.ReqIDKey), err)
		return err
	}
	return nil
}

func (service *SessionService) CheckVersion(ctx context.Context, login, token string,
	usersVersion uint32) (hasSession bool, err error) {
	hasSession, err = service.sessionStorage.CheckVersion(login, token, usersVersion)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to check version: %v", ctx.Value(requestId.ReqIDKey), err)
		return hasSession, err
	}
	return hasSession, nil
}

func (service *SessionService) GetVersion(ctx context.Context, login, token string) (version uint32, err error) {
	version, err = service.sessionStorage.GetVersion(login, token)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to get version: %v", ctx.Value(requestId.ReqIDKey), err)
		return version, err
	}
	return version, nil
}

func (service *SessionService) HasSession(ctx context.Context, login, token string) (err error) {
	err = service.sessionStorage.HasSession(login, token)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to has session: %v", ctx.Value(requestId.ReqIDKey), err)
		return err
	}
	return nil
}

func (service *SessionService) CheckAllUserSessionTokens(ctx context.Context, login string) error {
	err := service.sessionStorage.CheckAllUserSessionTokens(login)
	if err != nil {
		service.logger.Errorf("[reqid=%s] failed to check all user's session tokens: %v",
			ctx.Value(requestId.ReqIDKey), err)
		return err
	}
	return nil
}

<<<<<<< HEAD
<<<<<<< HEAD:internal/sessions/service/sessionservice.go
=======
func (service *SessionService) IsTokenValid(token *http.Cookie) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(service.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token: %w", myerrors.ErrNotAuthorised)
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token: %w", myerrors.ErrNotAuthorised)
	}

	_, ok = claims["Login"]
	if !ok {
		return nil, fmt.Errorf("invalid token: %w", myerrors.ErrNotAuthorised)
	}
	_, ok = claims["IsAdmin"]
	if !ok {
		return nil, fmt.Errorf("invalid token: %w", myerrors.ErrNotAuthorised)
	}

	_, ok = claims["Version"]
	if !ok {
		return nil, fmt.Errorf("invalid token: %w", myerrors.ErrNotAuthorised)
	}

	return claims, nil
}

func ValidateLogin(e string) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if emailRegex.MatchString(e) {
		return nil
	}
	return myerrors.ErrLoginIsNotValid
}

func ValidateUsername(username string) error {
	if len(username) >= 4 {
		return nil
	}
	return myerrors.ErrUsernameIsToShort
}

func ValidatePassword(password string) error {
	if len(password) >= 6 {
		return nil
	}
	return myerrors.ErrPasswordIsToShort
}

>>>>>>> a871897 (users done, waits for sessions):internal/service/sessionservice.go
=======
>>>>>>> d726255 (all microservices r done)
type customClaims struct {
	jwt.StandardClaims
	Login   string
	IsAdmin bool
	Version uint32
}

<<<<<<< HEAD:internal/sessions/service/sessionservice.go
func (service *SessionService) GenerateTokens(login string, isAdmin bool, version uint32) (tokenSigned string,
	err error) {
=======
func (service *SessionService) GenerateTokens(login string, isAdmin bool, version uint32) (tokenSigned string, err error) {
>>>>>>> a871897 (users done, waits for sessions):internal/service/sessionservice.go
	tokenCustomClaims := customClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    "NETrunnerFLIX",
		},
		Login:   login,
		IsAdmin: isAdmin,
		Version: version,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenCustomClaims)

	tokenSigned, err = token.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	return tokenSigned, nil
}
