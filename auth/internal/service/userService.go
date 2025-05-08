package service

import (
	"auth/internal/models"
	"auth/internal/repository"
	"context"
	"fmt"
)

type UserService interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	Login(ctx context.Context, user *models.User) (*models.User, error)
	AdminRegister(ctx context.Context, admin *models.Admin) (*models.Admin, error)
	AdminLogin(ctx context.Context, admin *models.Admin) (*models.Admin, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repository: repo}
}

func (u *userService) Register(ctx context.Context, user *models.User) (*models.User, error) {
	userResponse, err := u.repository.Register(ctx, user)
	if err != nil {
		fmt.Println("userService.register err: ", err)
		return nil, err
	}
	return userResponse, nil
}

func (u *userService) Login(ctx context.Context, user *models.User) (*models.User, error) {
	fmt.Println("userService.login user: ", user)
	return &models.User{}, nil
}

func (u *userService) AdminRegister(ctx context.Context, admin *models.Admin) (*models.Admin, error) {
	return &models.Admin{}, nil
}

func (u *userService) AdminLogin(ctx context.Context, admin *models.Admin) (*models.Admin, error) {
	return &models.Admin{}, nil
}
