package repository

import (
	"auth/internal/models"
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) CreateAdmin(ctx context.Context, amdin *models.Admin) (*models.Admin, error) {
	query := `INSERT INTO admins (username, password) VALUES ($1, $2) RETURNING id`
	err := r.DB.QueryRowContext(ctx, query, amdin.Username, amdin.Password).Scan(&amdin.ID)
	if err != nil {
		return nil, err
	}
	return amdin, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.DB.GetContext(ctx, &user, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAdminByUsername(ctx context.Context, username string) (*models.Admin, error) {
	var admin models.Admin
	err := r.DB.GetContext(ctx, &admin, "SELECT * FROM admins WHERE username=$1", username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("admin not found")
		}
		return nil, err
	}
	return &admin, nil
}
