package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/shoksin/marketplace-protos/proto/pbauth"
	"net/http"
)

type Handler struct {
	Client *Client
}

func NewHandler(client *Client) *Handler {
	return &Handler{Client: client}
}

func (h *Handler) Register(ctx *gin.Context) {
	var req struct {
		UserName string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.Client.Client.Register(context.Background(), &pbauth.RegisterRequest{
		Username: req.UserName,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
	}

	ctx.JSON(int(res.Status), res)
}

func (h *Handler) Login(ctx *gin.Context) {

}

func (h *Handler) AdminRegister(ctx *gin.Context) {

}

func (h *Handler) AdminLogin(ctx *gin.Context) {

}
