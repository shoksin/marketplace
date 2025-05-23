package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"order/internal/models"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, req *models.Order) (*models.Order, error)
}
type OrderService struct {
	repository OrderRepository
}

func NewOrderService(repository OrderRepository) *OrderService {
	return &OrderService{
		repository: repository,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error) {
	order.ID = uuid.New().String()
	resp, err := s.repository.CreateOrder(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("CreateOrder error: %w", err)
	}
	return resp, nil
}
