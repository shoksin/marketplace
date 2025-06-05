package order

import (
	"github.com/gin-gonic/gin"
	"github.com/shoksin/marketplace/api-gateway/internal/auth"
)

func SetupRoutes(r *gin.Engine, handler *Handler, authService *auth.Middleware) {
	order := r.Group("/order")
	order.Use(authService.ValidateUserToken())

	order.POST("/", handler.CreateOrder)
}
