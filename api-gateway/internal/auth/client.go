package auth

import (
	"github.com/shoksin/marketplace-protos/proto/pbauth"

	"google.golang.org/grpc"
)

type Client struct {
	Client pbauth.AuthServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		Client: pbauth.NewAuthServiceClient(conn),
	}
}
