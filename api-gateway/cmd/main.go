package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shoksin/marketplace/api-gateway/internal/auth"
	"github.com/shoksin/marketplace/api-gateway/internal/order"
	"github.com/shoksin/marketplace/api-gateway/internal/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

func mustNewClient(target string) *grpc.ClientConn {
	cp := grpc.ConnectParams{
		MinConnectTimeout: 5 * time.Second,
	}
	conn, err := grpc.NewClient(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(cp),
	)
	if err != nil {
		log.Fatalf("failed to connect to %s: %v", target, err)
	}
	return conn
}

func main() {
	authAddr := os.Getenv("AUTH_SERVICE") //например "auth-service:50051"
	productAddr := os.Getenv("PRODUCT_SERVICE")
	orderAddr := os.Getenv("ORDER_SERVICE")

	authConn := mustNewClient(authAddr)
	productConn := mustNewClient(productAddr)
	orderConn := mustNewClient(orderAddr)

	defer func() {
		if err := authConn.Close(); err != nil {
			log.Fatalf("Failed to close auth connection: %v", err)
		}
	}()

	defer func() {
		if err := productConn.Close(); err != nil {
			log.Fatalf("Failed to close product connection: %v", err)
		}
	}()

	defer func() {
		if err := orderConn.Close(); err != nil {
			log.Fatalf("Failed to close order connection: %v", err)
		}
	}()

	r := gin.Default()

	authClient := auth.NewClient(authConn)
	middleware := auth.NewMiddleware(authClient)
	authHandler := auth.NewHandler(authClient)
	auth.SetupRoutes(r, authHandler, middleware)

	productClient := product.NewClient(productConn)
	productHandler := product.NewHandler(productClient)
	product.SetupRoutes(r, productHandler, middleware)

	orderClient := order.NewClient(orderConn)
	orderHandler := order.NewHandler(orderClient)
	order.SetupRoutes(r, orderHandler, middleware)

	port := os.Getenv("PORT")

	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
