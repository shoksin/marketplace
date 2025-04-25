package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/shoksin/marketplace-protos/proto/pbauth"
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
	authorization, err := ctx.Cookie("token")
	if err != nil || authorization == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Missing authorization token",
		})
	}

	token := strings.TrimSpace(authorization)

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token is empty",
		})
	}

	res, err := m.client.Client.Validate(context.Background(), &pbauth.ValidateRequest{
		Token: token,
	})
	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": "invalid token",
		})
	}

	ctx.Set("userID", res.UserID)
	ctx.Next()
}
