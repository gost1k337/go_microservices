package server

import (
	"context"
	protos "currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	protos.UnimplementedCurrencyServer
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{log: l}
}

func (c *Currency) GetRate(ctx context.Context, r *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate:", "base", r.GetBase(), "quote", r.GetQuote())

	return &protos.RateResponse{Rate: 0.5}, nil
}
