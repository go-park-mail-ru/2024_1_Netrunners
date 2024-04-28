package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	helper "github.com/go-park-mail-ru/2024_1_Netrunners/cmd"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/films/api"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/films/repository"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/films/service"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
)

func main() {
	var (
		frontEndPort int
		backEndPort  int
		serverIP     string
	)
	flag.IntVar(&frontEndPort, "f-port", 8080, "front-end server port")
	flag.IntVar(&backEndPort, "b-port", 8020, "back-end server port")
	flag.StringVar(&serverIP, "ip", "94.139.247.246", "back-end server port")

	flag.Parse()

	err := helper.InitUploads()
	if err != nil {
		log.Fatal(err)
	}

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

	filmsStorage, err := repository.NewFilmsStorage(pool)
	if err != nil {
		log.Fatal(err)
	}

	filmService := service.NewFilmsService(filmsStorage, sugarLogger, "./uploads/films")

	s := grpc.NewServer()
	srv := api.NewFilmsServer(filmService, sugarLogger)
	session.RegisterFilmsServer(s, srv)

	listener, err := net.Listen("tcp", ":8020")
	if err != nil {
		log.Fatal(err)
	}

	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		if err := listener.Close(); err != nil {
			fmt.Printf("Server shutdown error: %v\n", err)
		}
	}()

	fmt.Printf("Starting server at %s%s\n", "localhost", fmt.Sprintf(":%d", backEndPort))

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

	<-stopped

	fmt.Println("Server stopped")
}
