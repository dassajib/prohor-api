package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userId uint) (string, error) {
	// define claims payload with store user id and expire within 15mins
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}
	// Create a new token with HMAC SHA256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// sign the token with the secret key from env
	return token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
}

func GenerateRefreshToken(userId uint) (string, error) {
	// define claims payload with store user id and expire within 15mins
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	// Create a new token with HMAC SHA256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// sign the token with the secret key from env
	return token.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
}
