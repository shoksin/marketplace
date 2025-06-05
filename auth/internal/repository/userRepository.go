package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/shoksin/marketplace/auth/internal/models"
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
	query := `INSERT INTO users (user_id, username, email, password, birthday, created_at) VALUES ($1, $2, $3, $4, $5, NOW()) RETURNING user_id, created_at`
	err := r.DB.QueryRowContext(ctx, query, user.ID, user.Username, user.Email, user.Password, user.Birthday).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) CreateAdmin(ctx context.Context, admin *models.Admin) (*models.Admin, error) {
	query := `INSERT INTO admins (admin_id, username, password, created_at) VALUES ($1, $2, $3, NOW()) RETURNING admin_id, created_at`
	err := r.DB.QueryRowContext(ctx, query, admin.ID, admin.Username, admin.Password).Scan(&admin.ID, &admin.CreatedAt)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := `SELECT user_id, username, password, email, birthday, created_at, updated_at, deleted_at FROM users WHERE email = $1`
	err := r.DB.GetContext(ctx, &user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	query := `SELECT user_id, username, password, email, birthday, created_at, updated_at, deleted_at FROM users WHERE username = $1`
	err := r.DB.GetContext(ctx, &user, query, username)
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

	err := r.DB.GetContext(ctx, &admin, `SELECT admin_id, username, password, created_at, updated_at, deleted_at FROM admins WHERE username = $1`, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("admin not found")
		}
		return nil, err
	}
	return &admin, nil
}
