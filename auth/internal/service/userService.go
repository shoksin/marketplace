package service

import (
	"auth/internal/dto"
	"auth/internal/models"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	CreateAdmin(ctx context.Context, admin *models.Admin) (*models.Admin, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetAdminByUsername(ctx context.Context, username string) (*models.Admin, error)
}

type TokenGenerator interface {
	GenerateUserToken(user *models.User) (string, error)
	GenerateAdminToken(admin *models.Admin) (string, error)
	ValidateToken(tokenString string, isAdmin bool) (*models.JWTClaims, error)
}

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type UserService struct {
	repository     UserRepository
	tokenGenerator TokenGenerator
	passwordHasher PasswordHasher
}

func NewUserService(repo UserRepository, tokenGenerator TokenGenerator, passwordHasher PasswordHasher) *UserService {
	return &UserService{
		repository:     repo,
		tokenGenerator: tokenGenerator,
		passwordHasher: passwordHasher,
	}
}

func (u *UserService) Register(ctx context.Context, user *models.User) (*models.User, error) {
	userExists, err := u.repository.GetUserByEmail(ctx, user.Email)
	if userExists != nil && userExists.Email != "" {
		return nil, fmt.Errorf("user with email %s already exists", user.Email)
	}

	userExists, err = u.repository.GetUserByUsername(ctx, user.Username)
	if userExists != nil && userExists.Username != "" {
		return nil, fmt.Errorf("user with username %s already exists", user.Username)
	}

	hashedPassword, err := u.passwordHasher.HashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashedPassword

	user.ID = uuid.New().String()

	resp, err := u.repository.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("register user error: %w", err)
	}
	return resp, nil
}

func (u *UserService) Login(ctx context.Context, user *models.User) (*dto.LoginResponse, error) {
	dbUser, err := u.repository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	if !u.passwordHasher.CheckPasswordHash(user.Password, dbUser.Password) {
		return nil, fmt.Errorf("invalid password")
	}
	user.ID = dbUser.ID
	token, err := u.tokenGenerator.GenerateUserToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &dto.LoginResponse{
		Token: token,
	}, nil
}

func (u *UserService) AdminRegister(ctx context.Context, admin *models.Admin) (*models.Admin, error) {
	adminExists, err := u.repository.GetAdminByUsername(ctx, admin.Username)
	if adminExists != nil && adminExists.Username != "" {
		return nil, fmt.Errorf("admin with username %s already exists", admin.Username)
	}

	hashedPassword, err := u.passwordHasher.HashPassword(admin.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	admin.Password = hashedPassword

	admin.ID = uuid.New().String()

	resp, err := u.repository.CreateAdmin(ctx, admin)
	if err != nil {
		return nil, fmt.Errorf("register admin error: %w", err)
	}
	return resp, nil
}

func (u *UserService) AdminLogin(ctx context.Context, admin *models.Admin) (*dto.LoginResponse, error) {
	dbAdmin, err := u.repository.GetAdminByUsername(ctx, admin.Username)
	if err != nil {
		return nil, fmt.Errorf("GetAdminByUsername error: %w", err)
	}

	if !u.passwordHasher.CheckPasswordHash(admin.Password, dbAdmin.Password) {
		return nil, fmt.Errorf("invalid password")
	}

	admin.ID = dbAdmin.ID

	token, err := u.tokenGenerator.GenerateAdminToken(admin)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &dto.LoginResponse{
		Token: token,
	}, nil
}
