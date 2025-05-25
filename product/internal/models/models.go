package models

import "time"

type Product struct {
	ID        string     `json:"id" db:"product_id"`
	Name      string     `json:"name" db:"name"`
	Price     float64    `json:"price" db:"price"`
	Stock     int64      `json:"stock" db:"stock"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type DecreaseStockRequest struct {
	ProductID string `json:"product_id" db:"product_id"`
	Quantity  int64  `json:"quantity" db:"quantity"`
}

type DecreaseStockResponse struct {
	ProductID string `json:"product_id" db:"product_id"`
}
