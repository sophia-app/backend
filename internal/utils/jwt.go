package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sophia-app/backend/internal/schemas"
)

// GetJWTPayload returns the JWT payload for the given user.
func GetJWTPayload(user schemas.User) jwt.MapClaims {
	return jwt.MapClaims{
		"userId": user.ID,
		"exp":    jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}
}

// GetJWTSecret returns the JWT secret from the environment variable.
func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}
