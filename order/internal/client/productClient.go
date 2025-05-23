package client

import (
	"context"
	"github.com/shoksin/marketplace-protos/proto/pbproduct"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type ProductServiceClient struct {
	Client pbproduct.ProductServiceClient
}

func NewProductServiceClient(target string) *ProductServiceClient {
	cp := grpc.ConnectParams{
		MinConnectTimeout: 5 * time.Second,
	}
	conn, err := grpc.NewClient(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(cp),
	)
	if err != nil {
		log.Fatalf("failed to connect to %s: %v", target, err)
	}
	return &ProductServiceClient{
		Client: pbproduct.NewProductServiceClient(conn),
	}
}

func (c *ProductServiceClient) FindOne(ctx context.Context, productID string) (*pbproduct.FindOneResponse, error) {
	req := &pbproduct.FindOneRequest{
		Id: productID,
	}
	return c.Client.FindOne(ctx, req)
}

func (c *ProductServiceClient) DecreaseStock(ctx context.Context, productID string, orderID string, quantity int64) (*pbproduct.DecreaseStockResponse, error) {
	req := &pbproduct.DecreaseStockRequest{
		Id:       productID,
		OrderID:  orderID,
		Quantity: quantity,
	}
	return c.Client.DecreaseStock(ctx, req)
}
