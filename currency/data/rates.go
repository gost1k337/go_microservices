package data

import (
	"encoding/json"
	"fmt"
	"github.com/gost1k337/go_microservices/currency/config"
	"github.com/hashicorp/go-hclog"
	"net/http"
)

type ExchangeRates struct {
	log   hclog.Logger
	cfg   *config.Config
	rates map[string]float32
}

func NewRates(log hclog.Logger, cfg *config.Config) (*ExchangeRates, error) {
	er := &ExchangeRates{
		log:   log,
		rates: map[string]float32{},
		cfg:   cfg,
	}

	err := er.GetRates()

	return er, err
}

func (er *ExchangeRates) GetRate(base, quote string) (float32, error) {
	br, ok := er.rates[base]
	if !ok {
		return 0, fmt.Errorf("rate not found for currency %s", base)
	}

	qr, ok := er.rates[quote]
	if !ok {
		return 0, fmt.Errorf("rate not found for currency %s", quote)
	}

	return qr / br, nil
}

func (er *ExchangeRates) GetRates() error {
	connString := fmt.Sprintf("http://api.exchangeratesapi.io/v1/latest?access_key=%s", er.cfg.ExchangeRateApiKey)
	resp, err := http.DefaultClient.Get(connString)
	if err != nil {
		return fmt.Errorf("error while getting rates: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected error code 200 got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	rr := &GetRatesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&rr); err != nil {
		return fmt.Errorf("parse response: %w", err)
	}

	for name, rate := range rr.Rates {
		er.rates[name] = rate
	}

	return nil
}
