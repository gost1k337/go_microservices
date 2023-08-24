package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/hashicorp/go-hclog"
	"net/http"
)

type Handlers struct {
	http *chi.Mux
	log  hclog.Logger
}

func New(logger hclog.Logger) *Handlers {
	h := &Handlers{
		log: logger,
	}

	r := chi.NewRouter()
	r.Use(middleware.DefaultLogger)
	r.Get("/", h.HelloWorld)

	h.http = r

	return h
}

func (h *Handlers) HTTP() http.Handler {
	return h.http
}

func (h *Handlers) HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
