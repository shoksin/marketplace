package handler

import (
	"auth/internal/models"
	"auth/internal/service"
	"context"
	"github.com/shoksin/marketplace-protos/proto/pbauth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
)

type AuthHandler struct {
	pbauth.UnimplementedAuthServiceServer
	service service.UserService
}

func NewAuthHandler(service service.UserService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(ctx context.Context, req *pbauth.RegisterRequest) (*pbauth.RegisterResponse, error) {
	user := &models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	res, err := h.service.Register(ctx, user)
	if err != nil {
		log.Println(err)
		return &pbauth.RegisterResponse{
			Error: err.Error(),
		}, nil
	}

	return &pbauth.RegisterResponse{
		Id:     res.ID,
		Status: http.StatusCreated,
	}, nil
}

func (h *AuthHandler) AdminRegister(ctx context.Context, req *pbauth.AdminRegisterRequest) (*pbauth.RegisterResponse, error) {
	admin := &models.Admin{
		Username: req.Username,
		Password: req.Password,
	}

	res, err := h.service.AdminRegister(ctx, admin)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbauth.RegisterResponse{
		Status: http.StatusCreated,
		Id:     res.ID,
	}, nil
}
func (h *AuthHandler) Login(ctx context.Context, req *pbauth.LoginRequest) (*pbauth.LoginResponse, error) {
	user := &models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := h.service.Login(ctx, user)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbauth.LoginResponse{
		Token: res.
	}, nil
}
func (h *AuthHandler) AdminLogin(context.Context, *pbauth.AdminLoginRequest) (*pbauth.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (h *AuthHandler) Validate(context.Context, *pbauth.ValidateRequest) (*pbauth.ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
