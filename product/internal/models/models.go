package models

import "time"

type Product struct {
	ID        int64      `json:"id" db:"product_id"`
	Name      string     `json:"name" db:"name"`
	Price     float64    `json:"price" db:"price"`
	Stock     int64      `json:"stock" db:"stock"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}
