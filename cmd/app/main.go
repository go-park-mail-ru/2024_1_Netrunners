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

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/handlers"
	mycache "github.com/go-park-mail-ru/2024_1_Netrunners/internal/repository/cache"
	mockdb "github.com/go-park-mail-ru/2024_1_Netrunners/internal/repository/mockDB"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/service"
)

func main() {
	data := []domain.FilmPreview{
		{
			Id:       "dfgea4ra424r4fw",
			Name:     "Film1",
			Duration: 3600,
		},
		{
			Id:       "fnuf7842huirn23",
			Name:     "Film2",
			Duration: 7200,
		},
	}

	cacheStorage := mycache.InitSessionStorage()
	authStorage := mockdb.InitUsersMockDB()
	filmsStorage := mockdb.InitFilmsMockDB()
	filmsStorage.AddFilm(data[0])
	filmsStorage.AddFilm(data[1])

	sessionService := service.InitSessionService(cacheStorage)
	authService := service.InitAuthService(authStorage)
	filmsService := service.InitFilmsService(filmsStorage)

	mainPageHandlers := handlers.InitMainPageHandlers()
	authPageHandlers := handlers.InitAuthPageHandlers(authService, sessionService)
	filmsPageHandlers := handlers.InitFilmsPageHandlers(filmsService)

	router := mux.NewRouter()

	router.HandleFunc("/", mainPageHandlers.GetIndex).Methods("GET")
	router.HandleFunc("/auth/login", authPageHandlers.Login).Methods("POST")
	router.HandleFunc("/auth/logout", authPageHandlers.Logout).Methods("POST")
	router.HandleFunc("/auth/signup", authPageHandlers.Signup).Methods("POST")
	router.HandleFunc("/auth/check", authPageHandlers.Check).Methods("POST")
	router.HandleFunc("/films", filmsPageHandlers.GetFilmsPreviews).Methods("GET")

	server := &http.Server{
		Handler: router,
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
