package handlers

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	proto "github.com/gost1k337/go_microservices/currency/protos/currency"
	"github.com/hashicorp/go-hclog"
	"net/http"
)

type Handlers struct {
	http *chi.Mux
	log  hclog.Logger
	cc   proto.CurrencyClient
}

func New(logger hclog.Logger, cc proto.CurrencyClient) *Handlers {
	h := &Handlers{
		log: logger,
		cc:  cc,
	}

	r := chi.NewRouter()
	r.Use(middleware.DefaultLogger)
	r.Get("/rates", h.GetRates)

	h.http = r

	return h
}

func (h *Handlers) HTTP() http.Handler {
	return h.http
}

func (h *Handlers) GetRates(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var req RateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("error while parsing request body", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	rr := &proto.RateRequest{
		Base:  proto.Currencies(proto.Currencies_value[req.Base]),
		Quote: proto.Currencies(proto.Currencies_value[req.Quote]),
	}
	resp, err := h.cc.GetRate(context.Background(), rr)
	if err != nil {
		h.log.Error("[Error] error getting new rate", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rate := RateResponse{
		Rate: resp.Rate,
	}
	if err := json.NewEncoder(w).Encode(rate); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
