package main

import (
	"auth/internal/handler"
	"auth/internal/initializer"
	"auth/internal/repository"
	"auth/internal/service"
	"auth/internal/utils"
	"github.com/shoksin/marketplace-protos/proto/pbauth"
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
