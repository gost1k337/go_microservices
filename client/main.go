package main

import (
	"fmt"
	"github.com/gost1k337/go_microservices/client_microservice/handlers"
	"github.com/gost1k337/go_microservices/client_microservice/server"
	protos "github.com/gost1k337/go_microservices/currency/protos/currency"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := hclog.Default()

	conn, err := grpc.Dial("localhost:9092", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("grpc: %w", err)
		os.Exit(1)
	}
	defer conn.Close()

	cc := protos.NewCurrencyClient(conn)

	h := handlers.New(log, cc)

	log.Info("Starting server on port: 8080")
	srv := server.NewServer(":8080", h.HTTP())
	fmt.Println(*srv)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err := <-srv.Notify():
		log.Error("app - Run - httpServer.Notify: %w", err)
	}

	// Graceful shutdown
	log.Info("Shutting down...")
	err = srv.Shutdown()
	if err != nil {
		log.Error("app - Run - httpServer.Shutdown: %w", err)
	}
}
