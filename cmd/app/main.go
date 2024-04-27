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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/handlers"
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
	flag.StringVar(&serverIP, "ip", "94.139.247.246", "back-end server port")

	flag.Parse()

	err := initUploads()
	if err != nil {
		log.Fatal(err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	sugarLogger := logger.Sugar()

	authConn, err := grpc.Dial(":8010", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	filmsConn, err := grpc.Dial(":8020", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	usersConn, err := grpc.Dial(":8030", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	csatConn, err := grpc.Dial(":8040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	filmsClient := session.NewFilmsClient(filmsConn)
	usersClient := session.NewUsersClient(usersConn)
	sessionClient := session.NewSessionsClient(authConn)
	csatClient := session.NewCsatClient(csatConn)

	middleware := middleware.NewMiddleware(sugarLogger, serverIP)
	authPageHandlers := handlers.NewAuthPageHandlers(&usersClient, &sessionClient, sugarLogger)
	usersPageHandlers := handlers.NewUserPageHandlers(&usersClient, &sessionClient, sugarLogger)
	filmsPageHandlers := handlers.NewFilmsPageHandlers(&filmsClient, sugarLogger)
	csatHandlers := handlers.NewCsatHandlers(&csatClient, sugarLogger)

	router := mux.NewRouter()

	router.HandleFunc("/auth/login", authPageHandlers.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/auth/logout", authPageHandlers.Logout).Methods("POST", "OPTIONS")
	router.HandleFunc("/auth/signup", authPageHandlers.Signup).Methods("POST", "OPTIONS")
	router.HandleFunc("/auth/check", authPageHandlers.Check).Methods("POST", "OPTIONS")

	router.HandleFunc("/films/all", filmsPageHandlers.GetAllFilmsPreviews).Methods("GET", "OPTIONS")
	router.HandleFunc("/films/{uuid}/data", filmsPageHandlers.GetFilmDataByUuid).Methods("GET", "OPTIONS")
	router.HandleFunc("/films/{uuid}/comments", filmsPageHandlers.GetAllFilmComments).Methods("GET", "OPTIONS")
	router.HandleFunc("/films/{uuid}/actors", filmsPageHandlers.GetActorsByFilm).Methods("GET", "OPTIONS")

	router.HandleFunc("/profile/{uuid}/data", usersPageHandlers.GetProfileData).Methods("GET", "OPTIONS")
	router.HandleFunc("/profile/{uuid}/edit", usersPageHandlers.ProfileEditByUuid).Methods("POST", "OPTIONS")
	router.HandleFunc("/profile/{uuid}/preview", usersPageHandlers.GetProfilePreview).Methods("GET", "OPTIONS")

	router.HandleFunc("/films",
		middleware.AuthMiddleware(filmsPageHandlers.GetAllFilmsPreviews)).Methods("GET", "OPTIONS")
	router.HandleFunc("/actors/{uuid}/data", filmsPageHandlers.GetActorByUuid).Methods("GET", "OPTIONS")

	router.HandleFunc("/csat/questions/get", csatHandlers.GetPageQuestions).Methods("GET", "OPTIONS")
	router.HandleFunc("/csat/stat/get", csatHandlers.GetStatisticsByPage).Methods("GET", "OPTIONS")
	router.HandleFunc("/csat/stat/add", csatHandlers.AddStatistics).Methods("POST", "OPTIONS")

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

func initUploads() error {
	storagePath := "./uploads/"
	_, err := os.Stat(storagePath)
	if err != nil {
		err = os.Mkdir(storagePath, 0755)
		if err != nil {
			return err
		}
		err = os.Mkdir(storagePath+"users/", 0755)
		if err != nil {
			return err
		}
		err = os.Mkdir(storagePath+"films/", 0755)
		if err != nil {
			return err
		}
	} else {
		storagePath = "./uploads/users/"
		_, err = os.Stat(storagePath)
		if err != nil {
			err = os.Mkdir(storagePath, 0755)
			if err != nil {
				return err
			}
		}

		storagePath := "./uploads/films/"
		_, err = os.Stat(storagePath)
		if err != nil {
			err = os.Mkdir(storagePath, 0755)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
