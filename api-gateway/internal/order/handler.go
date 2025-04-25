package order

import "github.com/gin-gonic/gin"

type Handler struct {
	client *Client
}

func NewHandler(client *Client) *Handler {
	return &Handler{
		client: client,
	}
}

func (h *Handler) CreateOrder(ctx *gin.Context) {

}
