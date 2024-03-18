package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/handlers"
	mycache "github.com/go-park-mail-ru/2024_1_Netrunners/internal/repository/cache"
	mockdb "github.com/go-park-mail-ru/2024_1_Netrunners/internal/repository/mockDB"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/repository/postgres"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
)

func CorsMiddleware(next http.Handler) http.Handler {
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

func main() {
	var (
		frontEndPort int
		backEndPort  int
	)
	flag.IntVar(&frontEndPort, "f-port", 8080, "front-end server port")
	flag.IntVar(&backEndPort, "b-port", 8081, "back-end server port")

	flag.Parse()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	sugarLogger := logger.Sugar()

	cacheStorage := mycache.InitSessionStorage()
	authStorage := mockdb.InitUsersMockDB()
	filmsStorage := mockdb.InitFilmsMockDB()
	actorsStorage, err := postgres.NewActorsStorage()
	if err != nil {
		log.Fatal(err)
	}

	sessionService := service.InitSessionService(cacheStorage)
	authService := service.InitAuthService(authStorage)
	_ = service.NewActorsService(actorsStorage, sugarLogger)
	filmsService := service.InitFilmsService(filmsStorage, "/root/2024_1_Netrunners/uploads")
	err = filmsService.AddSomeData()
	if err != nil {
		log.Fatal(err)
	}

	authPageHandlers := handlers.InitAuthPageHandlers(authService, sessionService)
	filmsPageHandlers := handlers.InitFilmsPageHandlers(filmsService)

	router := mux.NewRouter()

	router.HandleFunc("/auth/login", authPageHandlers.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/auth/logout", authPageHandlers.Logout).Methods("POST", "OPTIONS")
	router.HandleFunc("/auth/signup", authPageHandlers.Signup).Methods("POST", "OPTIONS")
	router.HandleFunc("/auth/check", authPageHandlers.Check).Methods("POST", "OPTIONS")
	router.HandleFunc("/films", filmsPageHandlers.GetFilmsPreviews).Methods("GET", "OPTIONS")

	router.Use(CorsMiddleware)

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", backEndPort),
	}

	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			fmt.Printf("Server shutdown error: %v\n", err)
		}
	}()

	fmt.Printf("Starting server at %s%s\n", "localhost", fmt.Sprintf(":%d", backEndPort))

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-stopped

	fmt.Println("Server stopped")
}
