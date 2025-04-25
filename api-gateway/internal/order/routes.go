package order

import (
	"api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, handler *Handler, middleware *auth.Middleware) {
	order := r.Group("/order")
}
