package main

import (
	"log"

	"github.com/Geovanny0401/microservice-go-grpc/order/config"
	"github.com/Geovanny0401/microservice-go-grpc/order/internal/adapters/db"
	"github.com/Geovanny0401/microservice-go-grpc/order/internal/adapters/grpc"
	"github.com/Geovanny0401/microservice-go-grpc/order/internal/adapters/payment"
	"github.com/Geovanny0401/microservice-go-grpc/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
