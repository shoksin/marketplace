package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"product/internal/models"
)

type ProductRepository struct {
	DB *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	query := `INSERT INTO products (product_id, name, price, stock, created_at) VALUES ($1, $2, $3, $4, NOW()) RETURNING product_id, created_at;`
	if err := r.DB.QueryRowContext(ctx, query, product.ID, product.Name, product.Price, product.Stock).Scan(&product.ID, &product.CreatedAt); err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) FindOneProductByID(ctx context.Context, ID string) (*models.Product, error) {
	var product models.Product
	query := `SELECT product_id, name, price, stock, created_at, updated_at, deleted_at FROM products WHERE product_id = $1;`
	if err := r.DB.GetContext(ctx, &product, query, ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) FindAllProducts(ctx context.Context) ([]*models.Product, error) {
	var products []*models.Product
	query := `SELECT product_id, name, price, stock, created_at, updated_at, deleted_at FROM products`
	rows, err := r.DB.QueryContext(ctx, query)
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (r *ProductRepository) DecreaseStock(ctx context.Context, productID string, stock int64) (*models.Product, error) {
	var product models.Product
	query := `UPDATE products SET stock = $1 WHERE product_id = $2 RETURNING product_id, stock;`
	if err := r.DB.GetContext(ctx, &product, query, stock, productID); err != nil {
		return &product, err
	}

	return &product, nil
}
