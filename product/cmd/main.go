package main

import (
	"github.com/shoksin/marketplace-protos/proto/pbproduct"
	"github.com/shoksin/marketplace/product/internal/handler"
	"github.com/shoksin/marketplace/product/internal/initializer"
	"github.com/shoksin/marketplace/product/internal/repository"
	"github.com/shoksin/marketplace/product/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	initializer.InitDB()
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
