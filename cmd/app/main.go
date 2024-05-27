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
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/handlers"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/metrics"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/middleware"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
)

func main() {
	var (
		frontEndPort int
		backEndPort  int
		serverIP     string
	)
	flag.IntVar(&frontEndPort, "f-port", 8080, "front-end server port")
	flag.IntVar(&backEndPort, "b-port", 8081, "back-end server port")
	flag.StringVar(&serverIP, "ip", "90.156.218.166", "back-end server port")

	flag.Parse()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	sugarLogger := logger.Sugar()

	authConn, err := grpc.Dial("sessions:8010", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	filmsConn, err := grpc.Dial("films:8020", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	usersConn, err := grpc.Dial("users:8030", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	httpMetrics := metrics.NewHttpMetrics()
	httpMetrics.Register()

	filmsClient := session.NewFilmsClient(filmsConn)
	usersClient := session.NewUsersClient(usersConn)
	sessionClient := session.NewSessionsClient(authConn)

	middleware := middleware.NewMiddleware(httpMetrics, sugarLogger, serverIP)
	authPageHandlers := handlers.NewAuthPageHandlers(&usersClient, &sessionClient, httpMetrics, sugarLogger)
	usersPageHandlers := handlers.NewUserPageHandlers(&usersClient, &sessionClient, httpMetrics, sugarLogger)
	filmsPageHandlers := handlers.NewFilmsPageHandlers(&filmsClient, httpMetrics, sugarLogger)

	// router := mux.NewRouter().Schemes("http").Subrouter()
	router := mux.NewRouter().Schemes("https").Subrouter()

	router.Handle("/metrics", promhttp.Handler())

	router.HandleFunc("/api/auth/login", authPageHandlers.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/auth/logout", authPageHandlers.Logout).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/auth/signup", authPageHandlers.Signup).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/auth/check", authPageHandlers.Check).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/films/all", filmsPageHandlers.GetAllFilmsPreviews).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/films/all_sub", filmsPageHandlers.GetFilmsPreviewsWithSub).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/films/{uuid}/data", filmsPageHandlers.GetFilmDataByUuid).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/films/{uuid}/actors", filmsPageHandlers.GetActorsByFilm).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/films/put_favorite", filmsPageHandlers.PutFavoriteFilm).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/films/remove_favorite", filmsPageHandlers.RemoveFavoriteFilm).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/films/{uuid}/all_favorite", filmsPageHandlers.GetAllFavoriteFilms).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/films/find/short", filmsPageHandlers.ShortSearch).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/films/find/long", filmsPageHandlers.LongSearch).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/films/genres/{uuid}/all", filmsPageHandlers.GetAllFilmsByGenre).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/films/genres/preview", filmsPageHandlers.GetAllGenres).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/films/add", filmsPageHandlers.AddFilm).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/films/top", filmsPageHandlers.GetTopFilms).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/films/comments/add", filmsPageHandlers.AddComment).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/films/comments/remove", filmsPageHandlers.RemoveComment).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/films/{uuid}/comments", filmsPageHandlers.GetAllFilmComments).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/profile/{uuid}/data", usersPageHandlers.GetProfileData).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/profile/{uuid}/edit", usersPageHandlers.ProfileEditByUuid).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/profile/{uuid}/preview", usersPageHandlers.GetProfilePreview).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/profile/{uuid}/subscriptions/check", usersPageHandlers.HasSubscription).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/profile/{uuid}/subscriptions/pay", usersPageHandlers.PaySubscription).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/subscriptions/get", usersPageHandlers.GetSubscriptions).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/films",
		middleware.AuthMiddleware(filmsPageHandlers.GetAllFilmsPreviews)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/actors/{uuid}/data", filmsPageHandlers.GetActorByUuid).Methods("GET", "OPTIONS")

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
