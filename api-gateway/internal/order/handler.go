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
		ProductID string `json:"product_id" binding:"required"`
		Quantity  int64  `json:"quantity" binding:"required"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		log.Printf("Request validation error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	val, exists := ctx.Get("user_id")
	if !exists {
		log.Printf("Request validation error: user_id = %v", val)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userID, ok := val.(string)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id format",
		})
		return
	}

	log.Println("userID:", userID)
	log.Printf("Product ID: %s\n\n\n\n\n", req.ProductID)

	res, err := h.client.client.CreateOrder(ctx, &pborder.CreateOrderRequest{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		UserID:    userID,
	})

	if err != nil || res == nil {
		log.Printf("Error when calling the gRPC service: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, res)
}
