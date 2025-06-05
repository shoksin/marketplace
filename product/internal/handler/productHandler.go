package handler

import (
	"context"
	"github.com/shoksin/marketplace-protos/proto/pbproduct"
	"github.com/shoksin/marketplace/product/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type ProductService interface {
	CreateProduct(context.Context, *models.Product) (*models.Product, error)
	FindOne(context.Context, string) (*models.Product, error)
	FindAll(context.Context) ([]*models.Product, error)
	DecreaseStock(context.Context, string, int64) (*models.Product, error)
}
type GrpcProductHandler struct {
	pbproduct.UnimplementedProductServiceServer
	service ProductService
}

func NewProductHandler(service ProductService) *GrpcProductHandler {
	return &GrpcProductHandler{service: service}
}

func (h *GrpcProductHandler) CreateProduct(ctx context.Context, req *pbproduct.CreateProductRequest) (*pbproduct.CreateProductResponse, error) {
	product := &models.Product{
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	}

	res, err := h.service.CreateProduct(ctx, product)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbproduct.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     res.ID,
	}, nil
}

func (h *GrpcProductHandler) FindOne(ctx context.Context, req *pbproduct.FindOneRequest) (*pbproduct.FindOneResponse, error) {
	res, err := h.service.FindOne(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if res == nil {
		return nil, nil
	}

	data := &pbproduct.FindOneData{
		Id:    res.ID,
		Name:  res.Name,
		Price: res.Price,
		Stock: res.Stock,
	}

	return &pbproduct.FindOneResponse{
		Data: data,
	}, nil
}

func (h *GrpcProductHandler) FindAll(ctx context.Context, req *pbproduct.FindAllRequest) (*pbproduct.FindAllResponse, error) {
	products, err := h.service.FindAll(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if products == nil {
		return nil, nil
	}

	var pbProducts []*pbproduct.FindOneData

	for _, product := range products {
		pbProducts = append(pbProducts, &pbproduct.FindOneData{
			Id:    product.ID,
			Name:  product.Name,
			Price: product.Price,
			Stock: product.Stock,
		})
	}

	return &pbproduct.FindAllResponse{
		Products: pbProducts,
	}, nil
}

func (h *GrpcProductHandler) DecreaseStock(ctx context.Context, req *pbproduct.DecreaseStockRequest) (*pbproduct.DecreaseStockResponse, error) {
	_, err := h.service.DecreaseStock(ctx, req.Id, req.Quantity)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pbproduct.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil

}
