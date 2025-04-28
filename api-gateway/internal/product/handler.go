package product

import (
	"context"
	"github.com/gin-gonic/gin"
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

func (h *Handler)CreateProduct(ctx *gin.Context) {
	var req struct {
		Name  string `json:"name" binding:"required"`
		Price int64  `json:"price" binding:"required"`
		Stock int64  `json:"stock" binding:"required"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		log.Printf("Request validation error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	res, err := h.client.client.CreateProduct(context.Background(), &pbproduct.CreateProductRequest{
		Name: req.Name,
		Price: req.Price,
		Stock: req.Stock,
	})

	if err != nil{

	}

	ctx.JSON(res.)

}
