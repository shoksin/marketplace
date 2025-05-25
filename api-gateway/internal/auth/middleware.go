package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/shoksin/marketplace-protos/proto/pbauth"
	"log"
	"net/http"
	"strings"
)

type Middleware struct {
	client *Client
}

func NewMiddleware(client *Client) *Middleware {
	return &Middleware{
		client: client,
	}
}

func (m *Middleware) ValidateUserToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization, err := ctx.Cookie("Authorization")
		if err != nil || authorization == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing authorization cookie",
			})
			return
		}

		token := strings.TrimSpace(authorization)

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token is empty",
			})
			return
		}

		res, err := m.client.Client.Validate(context.Background(), &pbauth.ValidateRequest{
			Token:   token,
			IsAdmin: false,
		})
		if err != nil || res.Status != http.StatusOK {
			log.Printf("Token validation error: %v", err)
			ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
				"error": "invalid token",
			})
			return
		}

		ctx.Set("user_id", res.ID)
		ctx.Next()
	}
}

func (m *Middleware) ValidateAdminToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization, err := ctx.Cookie("Authorization")
		if err != nil || authorization == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing authorization cookie",
			})
			return
		}

		token := strings.TrimSpace(authorization)

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token is empty",
			})
			return
		}

		res, err := m.client.Client.Validate(context.Background(), &pbauth.ValidateRequest{
			Token:   token,
			IsAdmin: true,
		})
		if err != nil || res.Status != http.StatusOK {
			log.Printf("Token validation error: %v", err)
			ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
				"error": "invalid token",
			})
			return
		}

		ctx.Set("admin_id", res.ID)
		ctx.Next()
	}
}
