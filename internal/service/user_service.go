package service

import (
	"errors"

	"github.com/dassajib/prohor-api/internal/model"
	"github.com/dassajib/prohor-api/internal/pkg/utils"
	"github.com/dassajib/prohor-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserService defines what user-related operations our application supports
// Think of this as the "public API" for user management
type UserService interface {
	Register(username, email, password, confirmPassword string) error // Handles user registration
	Login(email, password string) (string, string, error)             // Handles user login
}

// userService is the concrete implementation of UserService
type userService struct {
	repo repository.UserRepository // We depend on the repository to handle database operations
}

// NewUserService creates a ready-to-use user service
// We inject the repository dependency (database operations) when creating the service
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo} // Initialize with the provided repository
}

// Register handles new user registration with validation and password hashing
func (s *userService) Register(username, email, password, confirmPassword string) error {
	// 1. Check if passwords match
	if password != confirmPassword {
		return errors.New("passwords do not match")
	}

	// 2. Check if email already exists
	_, err := s.repo.FindByEmail(email)
	if err == nil { // No error means user was found
		return errors.New("email already registered")
	}

	// 3. Hash the password for secure storage
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// 4. Create new user object
	user := &model.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword), // Store the hashed version, never raw password!
	}

	// 5. Save to database
	return s.repo.Create(user)
}

// Login authenticates users and generates JWT tokens upon successful login
func (s *userService) Login(email, password string) (string, string, error) {
	// 1. Find user by email
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("user not found") // Don't reveal too much info (security best practice)
	}

	// 2. Compare provided password with stored hash
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid password") // Generic error message (security best practice)
	}

	// 3. Generate JWT tokens if credentials are valid
	accessToken, _ := utils.GenerateAccessToken(user.ID)   // Short-lived token
	refreshToken, _ := utils.GenerateRefreshToken(user.ID) // Long-lived token for getting new access tokens

	// 4. Return tokens
	return accessToken, refreshToken, nil
}
