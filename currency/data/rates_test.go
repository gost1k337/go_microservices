package data

import (
	"fmt"
	"github.com/gost1k337/go_microservices/currency/config"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestNewRates(t *testing.T) {
	err := godotenv.Load("/Users/gost1k/projects/go_microservices/currency/.env")
	if err != nil {
		t.Fatal(err)
	}

	cfg := &config.Config{}
	cfg.ExchangeRateApiKey = os.Getenv("EXCHANGE_RATE_API_KEY")

	tr, err := NewRates(hclog.Default(), cfg)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Rates %#v", tr.rates)
}
