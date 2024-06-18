package main

import (
	"microservice/config"
	"microservice/api"
	"microservice/pkg/grpc_client"
	"microservice/pkg/logger"
)

var (
	log        logger.Logger
	cfg        config.Config
	grpcClient *grpc_client.GrpcClient
)

func initDeps() {
	var err error
	cfg = config.Load()
	log = logger.New(cfg.LogLevel, "customer-api-gateway")

	grpcClient, err = grpc_client.New(cfg)
	if err != nil {
		log.Error("grpc dial error", logger.Error(err))
	}
}

func main() {
	initDeps()

	server := api.New(api.Config{
		Logger:     log,
		GrpcClient: grpcClient,
		Cfg:        cfg,
	})

	server.Run(cfg.HTTPPort)
}
