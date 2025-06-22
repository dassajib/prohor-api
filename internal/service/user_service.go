package service

import (
	"errors"

	"github.com/dassajib/prohor-api/internal/model"
	"github.com/dassajib/prohor-api/internal/pkg/utils"
	"github.com/dassajib/prohor-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserService defines the methods that user-related services must implement.
type UserService interface {
	Register(username, email, password, confirmPassword string) error
	Login(email, password string) (string, string, error)
}

// userService provides implementation of the UserService interface.
type userService struct {
	repo repository.UserRepository
}

// constructor func
// NewUserService creates and returns a new UserService instance.
func NewUserService(repo repository.UserRepository) userService {
	return userService{repo: repo}
}

// receiver func for registration logic, validation and db operation
func (s *userService) Register(username, email, password, confirmPassword string) error {
	// check password and confirm pass match or not
	if password != confirmPassword {
		return errors.New("Password do not match")
	}

	// check if the email already exists in the database
	_, err := s.repo.FindByEmail(email)
	if err == nil {
		return errors.New("Email already exist.")
	}

	// hash the user's password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// create new User model instance
	user := &model.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	// to save call repo Create method
	return s.repo.Create(user)
}

// receiver func for auth and generate token
func (s *userService) Login(email, password string) (string, string, error) {
	// find the user by email
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("User not found.")
	}

	// compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", "", errors.New("Wrong password.")
	}

	// generate access and refresh tokens
	accessToken, _ := utils.GenerateAccessToken(user.ID)
	refreshToken, _ := utils.GenerateRefreshToken(user.ID)

	// return generated tokens
	return accessToken, refreshToken, nil
}
