package models

import "time"

type Order struct {
	ID        string     `json:"order_id" db:"order_id"`
	ProductID string     `json:"product_id" db:"product_id"`
	Quantity  int64      `json:"quantity" db:"quantity"`
	UserID    string     `json:"user_id" db:"user_id"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}
