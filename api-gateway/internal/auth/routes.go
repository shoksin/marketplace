package auth

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine, handler *Handler) {
	auth := r.Group("/auth")
	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)

	admin := r.Group("/admin")
	admin.POST("/register", handler.AdminRegister)
	admin.POST("/login", handler.AdminLogin)
}
