package config

import (
	"fmt"
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort         string `env:"PORT"`
	ExchangeRateApiKey string `env:"EXCHANGE_RATE_API_KEY"`
	GrpcPort           string `env:"GRPC_PORT"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("load .env: %w", err)
	}

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return cfg, nil
}
