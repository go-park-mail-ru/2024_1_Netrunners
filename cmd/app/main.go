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

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

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
	flag.IntVar(&backEndPort, "b-port", 80, "back-end server port")

	flag.Parse()

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

	router.Use(corsMiddleware)
	// corsRouter := addCors(router, []string{fmt.Sprintf("http://localhost:%d/", frontEndPort)})

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
