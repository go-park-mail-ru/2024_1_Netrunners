package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/handlers"
	mycache "github.com/go-park-mail-ru/2024_1_Netrunners/internal/repository/cache"
	mockdb "github.com/go-park-mail-ru/2024_1_Netrunners/internal/repository/mockDB"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
	muxhandlers "github.com/gorilla/handlers"
)

func addCors(router *mux.Router, originNames []string) http.Handler {
	credentials := muxhandlers.AllowCredentials()
	methods := muxhandlers.AllowedMethods([]string{http.MethodGet, http.MethodPost})
	age := muxhandlers.MaxAge(3600)
	origins := muxhandlers.AllowedOrigins(originNames)
	handler := muxhandlers.CORS(credentials, methods, age, origins)(router)
	return handler
}

func main() {
	cacheStorage := mycache.InitSessionStorage()
	authStorage := mockdb.InitUsersMockDB()
	filmsStorage := mockdb.InitFilmsMockDB()

	sessionService := service.InitSessionService(cacheStorage)
	authService := service.InitAuthService(authStorage)
	filmsService := service.InitFilmsService(filmsStorage, "./uploads")
	err := filmsService.AddSomeData()
	if err != nil {
		log.Fatal(err)
	}

	authPageHandlers := handlers.InitAuthPageHandlers(authService, sessionService)
	filmsPageHandlers := handlers.InitFilmsPageHandlers(filmsService)

	router := mux.NewRouter()

	router.HandleFunc("/auth/login", authPageHandlers.Login).Methods("POST")
	router.HandleFunc("/auth/logout", authPageHandlers.Logout).Methods("POST")
	router.HandleFunc("/auth/signup", authPageHandlers.Signup).Methods("POST")
	router.HandleFunc("/auth/check", authPageHandlers.Check).Methods("POST")
	router.HandleFunc("/films", filmsPageHandlers.GetFilmsPreviews).Methods("GET")

	corsRouter := addCors(router, []string{"http://localhost:80/"})

	server := &http.Server{
		Handler: corsRouter,
		Addr:    ":80",
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

	fmt.Printf("Starting server at %s%s\n", "localhost", ":80")

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-stopped

	fmt.Println("Server stopped")
}
