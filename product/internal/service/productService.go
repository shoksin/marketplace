package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"product/internal/models"
)

type ProductRepository interface {
	CreateProduct(context.Context, *models.Product) (*models.Product, error)
	FindOneProductByID(context.Context, string) (*models.Product, error)
	FindAllProducts(context.Context) ([]*models.Product, error)
	DecreaseStock(context.Context, string, int64) (*models.Product, error)
}

type ProductService struct {
	repository ProductRepository
}

func NewProductService(repository ProductRepository) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	if product.Price < 0 || product.Stock < 1 || product.Name == "" {
		return nil, errors.New("invalid product data")
	}

	product.ID = uuid.New().String()

	return s.repository.CreateProduct(ctx, product)
}

func (s *ProductService) FindOne(ctx context.Context, id string) (*models.Product, error) {
	resp, err := s.repository.FindOneProductByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find product by id: %w", err)
	}
	return resp, nil
}

func (s *ProductService) FindAll(ctx context.Context) ([]*models.Product, error) {
	resp, err := s.repository.FindAllProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("find all products: %w", err)
	}
	return resp, nil
}

func (s *ProductService) DecreaseStock(ctx context.Context, productID string, quantity int64) (*models.Product, error) {
	log.Printf("DecreaseStock productID: %v, quantity: %v", productID, quantity)
	orderStat, err := s.repository.FindOneProductByID(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("find product by id: %w", err)
	}
	resp, err := s.repository.DecreaseStock(ctx, productID, orderStat.Stock-quantity)
	if err != nil {
		return nil, fmt.Errorf("decrease stock: %w", err)
	}
	return resp, nil
}
