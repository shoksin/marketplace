package order

import (
	"api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, handler *Handler, authService *auth.Middleware) {
	order := r.Group("/order")
	order.Use(authService.ValidateToken)

	order.POST("/", handler.CreateOrder)
}
