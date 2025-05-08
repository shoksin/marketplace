package utils

import (
	"auth/internal/models"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type JWTClaims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := JWTClaims{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "marketplace-auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretSigningKey := os.Getenv("SIGNING_KEY")
	if secretSigningKey == "" {
		return "", fmt.Errorf("SIGNING_KEY environment variable not set")
	}

	signedToken, err := token.SignedString([]byte(secretSigningKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateAdminToken(admin *models.Admin) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := JWTClaims{
		ID:       admin.ID,
		Username: admin.Username,
		Password: admin.Password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "marketplace-auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretSigningKey := os.Getenv("SIGNING_KEY")
	if secretSigningKey == "" {
		return "", fmt.Errorf("SIGNING_KEY environment variable not set")
	}

	signedToken, err := token.SignedString([]byte(secretSigningKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(tokenString string) (*JWTClaims, error) {
	secretSigningKey := os.Getenv("SIGNING_KEY")

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretSigningKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {

		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}
