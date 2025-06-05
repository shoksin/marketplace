package main

import (
	"github.com/shoksin/marketplace-protos/proto/pbauth"
	"github.com/shoksin/marketplace/auth/internal/handler"
	"github.com/shoksin/marketplace/auth/internal/initializer"
	"github.com/shoksin/marketplace/auth/internal/repository"
	"github.com/shoksin/marketplace/auth/internal/service"
	"github.com/shoksin/marketplace/auth/internal/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	initializer.InitDB()
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	userRepo := repository.NewUserRepository(initializer.DB)
	tokenGenerator := utils.NewTokenGenerator()
	passwordHasher := utils.NewPasswordHasher()
	userService := service.NewUserService(userRepo, tokenGenerator, passwordHasher)
	userHandler := handler.NewAuthHandler(userService, tokenGenerator)

	pbauth.RegisterAuthServiceServer(grpcServer, userHandler)
	log.Println("AuthService is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
