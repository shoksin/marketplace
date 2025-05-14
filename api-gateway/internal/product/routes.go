package product

import (
	"api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, handler *Handler, authService *auth.Middleware) {
	product := r.Group("/product")
	product.Use(authService.ValidateToken)
	product.POST("/", handler.CreateProduct)
	product.GET("/:id", handler.FindProduct)
	product.GET("/", handler.FindAllProducts)
}
