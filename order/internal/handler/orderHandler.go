package handler

import (
	"context"
	"github.com/shoksin/marketplace-protos/proto/pborder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"order/internal/client"
	"order/internal/models"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error)
}

type GrpcOrderHandler struct {
	pborder.UnimplementedOrderServiceServer
	service       OrderService
	productClient *client.ProductServiceClient
}

func NewGrpcOrderHandler(service OrderService, productClient *client.ProductServiceClient) *GrpcOrderHandler {
	return &GrpcOrderHandler{
		service:       service,
		productClient: productClient,
	}
}

func (h *GrpcOrderHandler) CreateOrder(ctx context.Context, req *pborder.CreateOrderRequest) (*pborder.CreateOrderResponse, error) {
	product, err := (h.productClient).FindOne(ctx, req.ProductID)
	if err != nil {
		log.Printf("FindOne error: %v\n", err)
		return &pborder.CreateOrderResponse{
			Status: http.StatusBadGateway,
		}, status.Error(codes.Internal, err.Error())
	}

	if product.Data.Stock < 1 {
		log.Printf("Product stock is less than 1\n")
		return &pborder.CreateOrderResponse{
			Status: http.StatusBadRequest,
		}, status.Error(codes.FailedPrecondition, "The stock must be greater than zero")
	}

	order := &models.Order{
		ProductID: product.Data.Id,
		Quantity:  req.Quantity,
		UserID:    req.UserID,
	}
	createOrderResp, err := h.service.CreateOrder(ctx, order)
	if err != nil {
		log.Println("Register error:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if _, err = h.productClient.DecreaseStock(ctx, createOrderResp.ProductID, createOrderResp.ID, createOrderResp.Quantity); err != nil {
		log.Println("Decrease Stock error:", err)
		return &pborder.CreateOrderResponse{
			Status: http.StatusBadRequest,
		}, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Order created ID: %v\n", createOrderResp.ID)

	return &pborder.CreateOrderResponse{
		Id: createOrderResp.ID,
	}, nil
}
