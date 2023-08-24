package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func NewServer(address string, handler http.Handler) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		Addr:         address,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	server := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: 30 * time.Second,
	}

	server.start()

	return server
}

func (s *Server) start() {
	err := s.server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
