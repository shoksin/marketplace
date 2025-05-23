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

func (m *Middleware) ValidateToken(ctx *gin.Context) {
	authorization, err := ctx.Cookie("Authorization")
	if err != nil || authorization == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Missing authorization cookie",
		})
		return
	}

	log.Printf("Authorization: %s", authorization)

	token := strings.TrimSpace(authorization)

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token is empty",
		})
		return
	}

	log.Printf("Token: %s", token)

	res, err := m.client.Client.Validate(context.Background(), &pbauth.ValidateRequest{
		Token: token,
	})
	if err != nil || res.Status != http.StatusOK {
		log.Printf("Token validation error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "invalid token",
		})
		return
	}

	log.Printf("Middleware ValidateToken user_id = %v", res.UserID)
	ctx.Set("user_id", res.UserID)
	ctx.Next()
}
