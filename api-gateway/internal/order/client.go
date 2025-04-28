package order

import (
	"github.com/shoksin/marketplace-protos/proto/pborder"
	"google.golang.org/grpc"
)

type Client struct {
	client pborder.OrderServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		client: pborder.NewOrderServiceClient(conn),
	}
}
