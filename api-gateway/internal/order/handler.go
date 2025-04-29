package order

import (
	"github.com/gin-gonic/gin"
	"github.com/shoksin/marketplace-protos/proto/pborder"
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
	}

	if err := ctx.ShouldBind(&req); err != nil {
		log.Printf("Request validation error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	val, exists := ctx.Get("user_id")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userId, ok := val.(int64)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id format",
		})
	}

	res, err := h.client.client.CreateOrder(ctx, &pborder.CreateOrderRequest{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		UserID:    userId,
	})

	if err != nil || res == nil {
		log.Printf("Error when calling the gRPC service: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, res)
}
