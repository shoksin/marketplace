package main

import (
	"api-gateway/internal/auth"
	"api-gateway/internal/order"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	authConn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to auth service: %v", err)
	}
	defer func() {
		if err := authConn.Close(); err != nil {
			log.Fatalf("Failed to close auth connection: %v", err)
		}
	}()

	productConn, err := grpc.NewClient(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to product service: %v", err)
	}
	defer func() {
		if err := productConn.Close(); err != nil {
			log.Fatalf("Failed to close product connection: %v", err)
		}
	}()

	orderConn, err := grpc.NewClient(":50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to order service: %v", err)
	}
	defer func() {
		if err := orderConn.Close(); err != nil {
			log.Fatalf("Failed to close order connection: %v", err)
		}
	}()

	r := gin.Default()

	authClient := auth.NewClient(authConn)
	authHandler := auth.NewHandler(authClient)
	auth.SetupRoutes(r, authHandler)

	middleware := auth.NewMiddleware(authClient)

	productClient := product.NewClient(productConn)
	productHandler := product.NewHandler(productClient, middleware)
	product.SetupRoutes(r, productHandler)

	orderClient := order.NewClient(orderConn)
	orderHandler := order.NewHandler(orderClient)
	order.SetupRoutes(r, orderHandler, middleware)
}
