package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"microservice_client/handlers"
	"microservice_client/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := hclog.Default()

	h := handlers.New(log)

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
	err := srv.Shutdown()
	if err != nil {
		log.Error("app - Run - httpServer.Shutdown: %w", err)
	}
}
