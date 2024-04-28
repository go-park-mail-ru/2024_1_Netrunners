package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/sessions/api"
	mycache "github.com/go-park-mail-ru/2024_1_Netrunners/internal/sessions/repository/cache"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/sessions/service"
)

func main() {
	var (
		frontEndPort int
		backEndPort  int
		serverIP     string
	)
	flag.IntVar(&frontEndPort, "f-port", 8080, "front-end server port")
	flag.IntVar(&backEndPort, "b-port", 8010, "back-end server port")
	flag.StringVar(&serverIP, "ip", "94.139.247.246", "back-end server port")

	flag.Parse()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	sugarLogger := logger.Sugar()

	cacheStorage := mycache.NewSessionStorage()
	if err != nil {
		log.Fatal(err)
	}

	sessionService := service.NewSessionService(cacheStorage, sugarLogger)

	s := grpc.NewServer()
	srv := api.NewSessionServer(sessionService, sugarLogger)
	session.RegisterSessionsServer(s, srv)

	listener, err := net.Listen("tcp", ":8010")
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
