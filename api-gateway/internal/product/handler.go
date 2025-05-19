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

func (h *Handler) CreateProduct(ctx *gin.Context) {
	var req struct {
		Name  string  `json:"name" binding:"required"`
		Price float64 `json:"price" binding:"required"`
		Stock int64   `json:"stock" binding:"required"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		log.Printf("Request validation error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.client.client.CreateProduct(context.Background(), &pbproduct.CreateProductRequest{
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	})

	if err != nil {
		log.Printf("Error when calling the gRPC service: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	if res == nil {
		log.Printf("nil response from CreateProduct")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "nil response from product service",
		})
		return
	}

	ctx.JSON(http.StatusOK, &pbproduct.CreateProductResponse{
		Status: http.StatusOK,
		Id:     res.Id,
	})

}

func (h *Handler) FindProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "cannot get product id",
		})
		return
	}

	product, err := h.client.client.FindOne(context.Background(), &pbproduct.FindOneRequest{
		Id: id,
	})

	if err != nil {
		log.Printf("Error when calling the gRPC service: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	if product == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "nil response from product service",
		})
		return
	}

	ctx.JSON(int(product.Status), &pbproduct.FindOneResponse{
		Data: product.Data,
	})
}

func (h *Handler) FindAllProducts(ctx *gin.Context) {
	products, err := h.client.client.FindAll(context.Background(), &pbproduct.FindAllRequest{})
	if err != nil || products == nil {
		log.Printf("Error when calling the gRPC service: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(int(products.Status), products)
}
