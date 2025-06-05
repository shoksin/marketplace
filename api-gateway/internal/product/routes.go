package product

import (
	"github.com/gin-gonic/gin"
	"github.com/shoksin/marketplace/api-gateway/internal/auth"
)

func SetupRoutes(r *gin.Engine, handler *Handler, authService *auth.Middleware) {
	product := r.Group("/product")
	product.Use(authService.ValidateAdminToken())
	product.POST("/", handler.CreateProduct)
	product.GET("/:id", handler.FindProduct)
	product.GET("/", handler.FindAllProducts)
}
