package repository

import (
	"auth/internal/models"
	"context"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	Login(ctx context.Context, email string, password string) (*models.User, error)
	AdminRegister(ctx context.Context, user *models.User) (*models.User, error)
	AdminLogin(ctx context.Context, username string, password string) (*models.User, error)
}

type userRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) Register(ctx context.Context, user *models.User) (*models.User, error) {
	return &models.User{}, nil
}

func (r *userRepository) Login(ctx context.Context, email string, password string) (*models.User, error) {
	return &models.User{}, nil
}

func (r *userRepository) AdminRegister(ctx context.Context, user *models.User) (*models.User, error) {
	return &models.User{}, nil
}

func (r *userRepository) AdminLogin(ctx context.Context, username, password string) (*models.User, error) {
	return &models.User{}, nil
}
