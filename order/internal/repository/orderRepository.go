package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/shoksin/marketplace/order/internal/models"
)

type OrderRepository struct {
	DB *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, orderData *models.Order) (*models.Order, error) {
	order := &models.Order{}
	query := `INSERT INTO orders(order_id, product_id, quantity, user_id, created_at) VALUES ($1, $2, $3, $4, NOW()) RETURNING order_id, product_id ,quantity, created_at`
	if err := r.DB.QueryRowContext(ctx, query, orderData.ID, orderData.ProductID, orderData.Quantity, orderData.UserID).Scan(&order.ID, &order.ProductID, &order.Quantity, &order.CreatedAt); err != nil {
		fmt.Printf("Error inserting order: %v\n", err)
		return nil, err
	}
	return order, nil
}
