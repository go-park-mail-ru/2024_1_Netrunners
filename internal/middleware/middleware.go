package middleware

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/handlers"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
	"go.uber.org/zap"

	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
)

type Middleware struct {
	authService    *service.AuthService
	sessionService *service.SessionService
	logger         *zap.SugaredLogger
}

func NewMiddleware(authService *service.AuthService,
	sessionService *service.SessionService, logger *zap.SugaredLogger) *Middleware {
	return &Middleware{
		authService:    authService,
		sessionService: sessionService,
		logger:         logger,
	}
}

func (middlewareHandlers *Middleware) CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://94.139.247.246:8080")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, "+
			"Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (middlewareHandlers *Middleware) PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				middlewareHandlers.logger.Fatalf("panic raised from %v: %v", r.URL, err)
				err = handlers.WriteError(w, myerrors.ErrInternalServerError)
				if err != nil {
					middlewareHandlers.logger.Errorf("error at writing response: %v\n", err)
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (middleware *Middleware) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// here will be our new check function

		next.ServeHTTP(w, r)
	}
}
