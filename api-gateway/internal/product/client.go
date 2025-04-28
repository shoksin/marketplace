package product

import (
	"github.com/shoksin/marketplace-protos/proto/pbproduct"
	"google.golang.org/grpc"
)

type Client struct {
	client pbproduct.ProductServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		client: pbproduct.NewProductServiceClient(conn),
	}
}
