package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/shoksin/marketplace/auth/internal/models"
	"log"
	"os"
	"time"
)

type JWTGenerator struct {
}

func NewTokenGenerator() *JWTGenerator {
	return &JWTGenerator{}
}

func (tkGen *JWTGenerator) GenerateUserToken(user *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := models.JWTClaims{
		ID:    user.ID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "marketplace-auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretSigningKey := os.Getenv("USER_SECRET_KEY")
	if secretSigningKey == "" {
		return "", fmt.Errorf("USER_SECRET_KEY environment variable not set")
	}

	signedToken, err := token.SignedString([]byte(secretSigningKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (tkGen *JWTGenerator) GenerateAdminToken(admin *models.Admin) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := models.JWTClaims{
		ID:       admin.ID,
		Username: admin.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "marketplace-auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretSigningKey := os.Getenv("ADMIN_SECRET_KEY")
	if secretSigningKey == "" {
		return "", fmt.Errorf("ADMIN_SECRET_KEY environment variable not set")
	}

	signedToken, err := token.SignedString([]byte(secretSigningKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (tkGen *JWTGenerator) ValidateToken(tokenString string, isAdmin bool) (*models.JWTClaims, error) {
	var secretSigningKey string
	if !isAdmin {
		secretSigningKey = os.Getenv("USER_SECRET_KEY")
	} else {
		secretSigningKey = os.Getenv("ADMIN_SECRET_KEY")
	}

	if secretSigningKey == "" {
		return nil, fmt.Errorf("USER_SECRET_KEY or ADMIN_SECRET_KEY environment variable not set")
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretSigningKey), nil
	})
	if err != nil {
		log.Println("ParseWithClaims error:", err)
		return nil, err
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {

		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}
