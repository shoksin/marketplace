package main

import (
	"github.com/shoksin/marketplace-protos/proto/pbproduct"
	"google.golang.org/grpc"
	"log"
	"net"
	"product/internal/handler"
	"product/internal/initializer"
	"product/internal/repository"
	"product/internal/service"
)

func init() {
	initializer.InitDB()
	//initializer.LoadConfig()
}

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	productRepo := repository.NewProductRepository(initializer.DB)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	pbproduct.RegisterProductServiceServer(grpcServer, productHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
