package models

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type User struct {
	ID        string     `json:"id" db:"user_id"`
	Username  string     `json:"username" db:"username"`
	Password  string     `json:"password" db:"password"`
	Email     string     `json:"email" db:"email"`
	Birthday  string     `json:"birthday" db:"birthday"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type Admin struct {
	ID        string     `json:"id" db:"admin_id"`
	Username  string     `json:"username" db:"username"`
	Password  string     `json:"password" db:"password"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type JWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
