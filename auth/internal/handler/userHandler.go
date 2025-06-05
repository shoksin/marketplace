package handler

import (
	"context"
	"github.com/shoksin/marketplace-protos/proto/pbauth"
	"github.com/shoksin/marketplace/auth/internal/dto"
	"github.com/shoksin/marketplace/auth/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
)

type UserService interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	Login(ctx context.Context, user *models.User) (*dto.LoginResponse, error)
	AdminRegister(ctx context.Context, admin *models.Admin) (*models.Admin, error)
	AdminLogin(ctx context.Context, admin *models.Admin) (*dto.LoginResponse, error)
}

type TokenGenerator interface {
	GenerateUserToken(user *models.User) (string, error)
	GenerateAdminToken(admin *models.Admin) (string, error)
	ValidateToken(tokenString string, isAdmin bool) (*models.JWTClaims, error)
}

type GrpcAuthHandler struct {
	pbauth.UnimplementedAuthServiceServer
	service      UserService
	jwtGenerator TokenGenerator
}

func NewAuthHandler(service UserService, generator TokenGenerator) *GrpcAuthHandler {
	return &GrpcAuthHandler{service: service, jwtGenerator: generator}
}

func (h *GrpcAuthHandler) Register(ctx context.Context, req *pbauth.RegisterRequest) (*pbauth.RegisterResponse, error) {
	user := &models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Birthday: req.Birthday,
	}

	res, err := h.service.Register(ctx, user)
	if err != nil {
		log.Println("Register error:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbauth.RegisterResponse{
		Id:     res.ID,
		Status: http.StatusCreated,
	}, nil
}

func (h *GrpcAuthHandler) AdminRegister(ctx context.Context, req *pbauth.AdminRegisterRequest) (*pbauth.RegisterResponse, error) {
	admin := &models.Admin{
		Username: req.Username,
		Password: req.Password,
	}

	res, err := h.service.AdminRegister(ctx, admin)
	if err != nil {
		log.Println("AdminRegister error:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbauth.RegisterResponse{
		Status: http.StatusCreated,
		Id:     res.ID,
	}, nil
}
func (h *GrpcAuthHandler) Login(ctx context.Context, req *pbauth.LoginRequest) (*pbauth.LoginResponse, error) {
	loginReq := &models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := h.service.Login(ctx, loginReq)
	if err != nil {
		log.Println("Login error:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbauth.LoginResponse{
		Token:  res.Token,
		Status: http.StatusOK,
	}, nil
}
func (h *GrpcAuthHandler) AdminLogin(ctx context.Context, req *pbauth.AdminLoginRequest) (*pbauth.LoginResponse, error) {
	loginReq := &models.Admin{
		Username: req.Username,
		Password: req.Password,
	}

	res, err := h.service.AdminLogin(ctx, loginReq)
	if err != nil {
		log.Println("AdminLogin error:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbauth.LoginResponse{
		Token:  res.Token,
		Status: http.StatusOK,
	}, nil
}
func (h *GrpcAuthHandler) Validate(ctx context.Context, req *pbauth.ValidateRequest) (*pbauth.ValidateResponse, error) {
	claims, err := h.jwtGenerator.ValidateToken(req.Token, req.IsAdmin)
	if err != nil || claims == nil {
		log.Println("Validate error:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbauth.ValidateResponse{
		ID:     claims.ID,
		Status: http.StatusOK,
	}, nil
}
