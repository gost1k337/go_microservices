package main

import (
	"fmt"
	"github.com/gost1k337/go_microservices/currency/config"
	"github.com/gost1k337/go_microservices/currency/data"
	protos "github.com/gost1k337/go_microservices/currency/protos/currency"
	"github.com/gost1k337/go_microservices/currency/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	logger := hclog.Default()

	gs := grpc.NewServer()
	rates, err := data.NewRates(logger, cfg)
	if err != nil {
		logger.Error("error while getting rates", err)
	}
	cs := server.NewCurrency(rates, logger)

	protos.RegisterCurrencyServer(gs, cs)
	reflection.Register(gs)

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GrpcPort))
	if err != nil {
		logger.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	logger.Info(fmt.Sprintf("GRPC server started listening on port %s...", cfg.GrpcPort))
	_ = gs.Serve(l)
}
