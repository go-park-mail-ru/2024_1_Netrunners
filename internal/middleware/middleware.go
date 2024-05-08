package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/handlers"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/metrics"
	reqid "github.com/go-park-mail-ru/2024_1_Netrunners/internal/requestId"
)

type Middleware struct {
	metrics  *metrics.HttpMetrics
	logger   *zap.SugaredLogger
	serverIP string
}

func NewMiddleware(metrics *metrics.HttpMetrics, logger *zap.SugaredLogger, serverIP string) *Middleware {
	return &Middleware{
		metrics:  metrics,
		logger:   logger,
		serverIP: serverIP,
	}
}

func (middlewareHandlers *Middleware) CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", fmt.Sprintf("http://%s:8080", middlewareHandlers.serverIP))
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
				err = handlers.WriteError(w, r, middlewareHandlers.metrics, myerrors.ErrInternalServerError)
				if err != nil {
					middlewareHandlers.logger.Errorf("error at writing response: %v\n", err)
				}

				middlewareHandlers.metrics.IncRequestsTotal(r.URL.Path, r.Method, 500)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (middlewareHandlers *Middleware) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// here will be our new check function

		next.ServeHTTP(w, r)
	}
}

func (middlewareHandlers *Middleware) AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqId := reqid.GenerateRequestID()
		ctx := r.Context()
		ctx = context.WithValue(ctx, reqid.ReqIDKey, reqId)
		middlewareHandlers.logger.Info("request accessLog", "path", r.URL.Path)
		start := time.Now()
		next.ServeHTTP(w, r.WithContext(ctx))
		middlewareHandlers.logger.Info(fmt.Sprintf("requestProcessed reqid[%s], method[%s], URLPath[%s], "+
			"time = [%s];",
			reqId, r.Method, r.URL.Path, time.Since(start)))
	})
}
