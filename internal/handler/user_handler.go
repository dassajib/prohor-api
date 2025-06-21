package handler

import (
	"net/http"

	"github.com/dassajib/prohor-api/internal/service"
	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests related to users
type UserHandler struct {
	service service.UserService // We depend on the service layer for business logic
}

// NewUserHandler creates a new UserHandler instance
// We inject the UserService dependency when creating the handler
func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Register handles user registration HTTP requests
func (h *UserHandler) Register(c *gin.Context) {
	// 1. Define a struct to capture JSON request body
	var body struct {
		Username        string `json:"username"`         // Expected in JSON payload
		Email           string `json:"email"`            // Expected in JSON payload
		Password        string `json:"password"`         // Expected in JSON payload
		ConfirmPassword string `json:"confirm_password"` // Expected in JSON payload
	}

	// 2. Bind JSON request body to our struct
	if err := c.ShouldBindJSON(&body); err != nil {
		// Return 400 Bad Request if JSON is invalid
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// 3. Call service layer to handle registration
	err := h.service.Register(body.Username, body.Email, body.Password, body.ConfirmPassword)
	if err != nil {
		// Return 400 with error message if registration fails
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 4. Return 201 Created on success
	c.JSON(http.StatusCreated, gin.H{"message": "registration successful"})
}

// Login handles user authentication HTTP requests
func (h *UserHandler) Login(c *gin.Context) {
	// 1. Define a struct to capture login credentials
	var body struct {
		Email    string `json:"email"`    // Expected in JSON payload
		Password string `json:"password"` // Expected in JSON payload
	}

	// 2. Bind JSON request body to our struct
	if err := c.ShouldBindJSON(&body); err != nil {
		// Return 400 Bad Request if JSON is invalid
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// 3. Call service layer to authenticate user and generate tokens
	token, refreshToken, err := h.service.Login(body.Email, body.Password)
	if err != nil {
		// Return 401 Unauthorized if login fails
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 4. Return 200 OK with JWT tokens on success
	c.JSON(http.StatusOK, gin.H{
		"access_token":  token,        // Short-lived token for API access
		"refresh_token": refreshToken, // Long-lived token for getting new access tokens
	})
}
