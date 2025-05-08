package main

import (
	"auth/internal/handler"
	"auth/internal/initializer"
	"auth/internal/repository"
	"auth/internal/service"
	"github.com/shoksin/marketplace-protos/proto/pbauth"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	initializer.LoadEnv()
	initializer.InitDB()
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	db := initializer.DB
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewAuthHandler(userService)

	pbauth.RegisterAuthServiceServer(grpcServer, userHandler)
	log.Println("AuthService is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
