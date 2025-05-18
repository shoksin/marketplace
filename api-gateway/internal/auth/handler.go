package auth

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shoksin/marketplace-protos/proto/pbauth"
	"log"
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
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		Birthday string `json:"birthday" binding:"required"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		log.Printf("Request validation error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.Client.Client.Register(context.Background(), &pbauth.RegisterRequest{
		Username: req.UserName,
		Email:    req.Email,
		Password: req.Password,
		Birthday: req.Birthday,
	})

	if err != nil {
		log.Printf("Error when calling the gRPC service: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(int(res.Status), res)
}

func (h *Handler) Login(ctx *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Printf("Validation error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.Client.Client.Login(context.Background(), &pbauth.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		log.Printf("Error when calling the gRPC service: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.SetCookie("token", res.Token, 3600*24*30, "", "", false, true)

	ctx.JSON(int(res.Status), res)
}

func (h *Handler) AdminRegister(ctx *gin.Context) {
	var req struct {
		UserName string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		log.Printf("Request validation error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.Client.Client.AdminRegister(context.Background(), &pbauth.AdminRegisterRequest{
		Username: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		log.Printf("Error when calling the gRPC service: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(int(res.Status), res)
}

func (h *Handler) AdminLogin(ctx *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		log.Printf("Request validation error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.Client.Client.AdminLogin(context.Background(), &pbauth.AdminLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		log.Printf("Error when calling the gRPC service: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(int(res.Status), res)
}
