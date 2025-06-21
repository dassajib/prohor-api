package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateAccessToken creates a short-lived JWT token for authentication
// Parameters:
//   - userID: The unique identifier of the user
//
// Returns:
//   - string: The generated JWT token
//   - error: Any error that occurred during token generation
func GenerateAccessToken(userID uint) (string, error) {
	// Define the token claims (payload)
	claims := jwt.MapClaims{
		"user_id": userID,                                  // Store user ID in the token
		"exp":     time.Now().Add(15 * time.Minute).Unix(), // Token expires in 15 minutes
	}

	// Create a new token with HMAC SHA256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret key from environment variables
	// ACCESS_SECRET should be a long, random string (keep it secret!)
	return token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
}

// GenerateRefreshToken creates a long-lived JWT token for obtaining new access tokens
// Parameters:
//   - userID: The unique identifier of the user
//
// Returns:
//   - string: The generated refresh token
//   - error: Any error that occurred during token generation
func GenerateRefreshToken(userID uint) (string, error) {
	// Define the token claims (payload)
	claims := jwt.MapClaims{
		"user_id": userID,                                    // Store user ID in the token
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(), // Token expires in 7 days
	}

	// Create a new token with HMAC SHA256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our refresh secret key from environment variables
	// REFRESH_SECRET should be different from ACCESS_SECRET
	return token.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
}
