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
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/handlers"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/middleware"
	mycache "github.com/go-park-mail-ru/2024_1_Netrunners/internal/repository/cache"
	database "github.com/go-park-mail-ru/2024_1_Netrunners/internal/repository/postgres"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
)

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

	pool, err := pgxpool.New(context.Background(), fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"postgres",
		"root1234",
		"netrunnerflix",
	))
	if err != nil {
		log.Fatal(err)
	}

	cacheStorage := mycache.NewSessionStorage()
	authStorage, err := database.NewUsersStorage(pool)
	if err != nil {
		log.Fatal(err)
	}
	filmsStorage, err := database.NewFilmsStorage(pool)
	if err != nil {
		log.Fatal(err)
	}
	actorsStorage, err := database.NewActorsStorage(pool)
	if err != nil {
		log.Fatal(err)
	}

	sessionService := service.NewSessionService(cacheStorage, sugarLogger)
	authService := service.NewAuthService(authStorage, sugarLogger)
	actorsService := service.NewActorsService(actorsStorage, sugarLogger)
	filmsService := service.NewFilmsService(filmsStorage, sugarLogger, "/root/2024_1_Netrunners/uploads")

	middleware := middleware.NewMiddleware(authService, sessionService, sugarLogger)
	authPageHandlers := handlers.NewAuthPageHandlers(authService, sessionService, sugarLogger)
	filmsPageHandlers := handlers.NewFilmsPageHandlers(filmsService, sugarLogger)
	usersPageHandlers := handlers.NewUserPageHandlers(authService, sessionService, sugarLogger)
	actorsPageHandlers := handlers.NewActorsHandlers(actorsService, sugarLogger)

	// err = filmsService.AddSomeData()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/auth/login", authPageHandlers.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/auth/logout", authPageHandlers.Logout).Methods("POST", "OPTIONS")
	router.HandleFunc("/auth/signup", authPageHandlers.Signup).Methods("POST", "OPTIONS")
	router.HandleFunc("/auth/check", authPageHandlers.Check).Methods("POST", "OPTIONS")

	router.HandleFunc("/films/all", filmsPageHandlers.GetAllFilmsPreviews).Methods("GET", "OPTIONS")
	router.HandleFunc("/films/{uuid}/data", filmsPageHandlers.GetFilmDataByUuid).Methods("GET", "OPTIONS")
	router.HandleFunc("/films/{uuid}/comments", filmsPageHandlers.GetAllFilmComments).Methods("GET", "OPTIONS")
	router.HandleFunc("/films/{uuid}/actors", filmsPageHandlers.GetAllFilmActors).Methods("GET", "OPTIONS")
	router.HandleFunc("/films/add", filmsPageHandlers.AddFilm).Methods("POST", "OPTIONS")

	router.HandleFunc("/profile/{uuid}/data", usersPageHandlers.GetProfileData).Methods("GET", "OPTIONS")
	router.HandleFunc("/profile/{uuid}/edit", usersPageHandlers.ProfileEditByUuid).Methods("POST", "OPTIONS")
	router.HandleFunc("/profile/{uuid}/preview", usersPageHandlers.GetProfilePreview).Methods("GET", "OPTIONS")

	router.HandleFunc("/films",
		middleware.AuthMiddleware(filmsPageHandlers.GetAllFilmsPreviews)).Methods("GET", "OPTIONS")
	router.HandleFunc("/actors/{uuid}/data", actorsPageHandlers.GetActorByUuid).Methods("GET", "OPTIONS")

	router.Use(middleware.CorsMiddleware)
	router.Use(middleware.PanicMiddleware)
	router.Use(middleware.AccessLogMiddleware)

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
