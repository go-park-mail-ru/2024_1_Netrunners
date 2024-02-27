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
)

func main() {
	mainPageHandlers := handlers.InitMainPageHandlers()

	router := mux.NewRouter()

	router.HandleFunc("/", mainPageHandlers.GetIndex).Methods("GET")

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
