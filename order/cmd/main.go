package main

import (
	"github.com/shoksin/marketplace-protos/proto/pborder"
	"github.com/shoksin/marketplace/order/internal/client"
	"github.com/shoksin/marketplace/order/internal/handler"
	"github.com/shoksin/marketplace/order/internal/initializer"
	"github.com/shoksin/marketplace/order/internal/repository"
	"github.com/shoksin/marketplace/order/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func init() {
	initializer.InitDB()
	//initializer.LoadConfig()
}

func main() {
	productAddr := os.Getenv("PRODUCT_SERVICE")
	productClient := client.NewProductServiceClient(productAddr)

	db := initializer.DB
	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	orderHandler := handler.NewGrpcOrderHandler(orderService, productClient)

	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pborder.RegisterOrderServiceServer(grpcServer, orderHandler)

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
