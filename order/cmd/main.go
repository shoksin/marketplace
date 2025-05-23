package main

import (
	"github.com/shoksin/marketplace-protos/proto/pborder"
	"google.golang.org/grpc"
	"log"
	"net"
	"order/internal/client"
	"order/internal/handler"
	"order/internal/initializer"
	"order/internal/repository"
	"order/internal/service"
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
