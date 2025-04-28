package order

import (
	"github.com/gin-gonic/gin"
	"github.com/shoksin/marketplace-protos/proto/pborder"
	"github.com/shoksin/marketplace-protos/proto/pbproduct"
	"log"
	"net/http"
)

type Handler struct {
	client *Client
}

func NewHandler(client *Client) *Handler {
	return &Handler{
		client: client,
	}
}

func (h *Handler) CreateOrder(ctx *gin.Context) {
	var req struct {
		ProductID int64 `json:"product_id" binding:"required"`
		Quantity  int64 `json:"quantity" binding:"required"`
		UserID    int64 `json:"user_id" binding:"required"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		log.Printf("Request validation error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	res, err := h.client.client.CreateOrder(ctx, &pborder.CreateOrderRequest{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		UserID:    req.UserID,
	})

	if err != nil {
		log.Printf("Error when calling the gRPC service: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}
}
