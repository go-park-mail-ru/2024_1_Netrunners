package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/metrics"
	session "github.com/go-park-mail-ru/2024_1_Netrunners/internal/session/proto"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/sessions/api"
	mycache "github.com/go-park-mail-ru/2024_1_Netrunners/internal/sessions/repository/cache"
	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/sessions/service"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	// init metrics handler
	grpcMetrics := metrics.InitGrpcMetrics("auth")
	grpcMetrics.Register()

	router := mux.NewRouter()

	router.Handle("/metrics", promhttp.Handler())
	metricsServer := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", backEndPort+1),
	}
	if err := metricsServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	fmt.Printf("Starting metrics server at %s%s\n", "localhost", fmt.Sprintf(":%d", backEndPort+1))

	sessionService := service.NewSessionService(cacheStorage, grpcMetrics, sugarLogger)

	// init grpc server
	s := grpc.NewServer()
	srv := api.NewSessionServer(sessionService, sugarLogger)
	session.RegisterSessionsServer(s, srv)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", backEndPort))
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
