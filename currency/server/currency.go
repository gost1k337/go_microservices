package server

import (
	"context"
	"github.com/gost1k337/go_microservices/currency/data"
	protos "github.com/gost1k337/go_microservices/currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	protos.UnimplementedCurrencyServer
	rates *data.ExchangeRates
	log   hclog.Logger
}

func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{rates: r, log: l}
}

func (c *Currency) GetRate(ctx context.Context, r *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate:", "base", r.GetBase(), "quote", r.GetQuote())
	rate, err := c.rates.GetRate(r.GetBase().String(), r.GetQuote().String())
	if err != nil {
		return nil, err
	}
	return &protos.RateResponse{Rate: rate}, nil
}
