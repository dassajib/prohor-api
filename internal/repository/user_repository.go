package repository

import (
	"github.com/dassajib/prohor-api/internal/model"
	"gorm.io/gorm"
)

// UserRepository is like a contract that defines what user-related database operations we support
// Think of it as a menu at a restaurant - it shows what you can order
type UserRepository interface {
	Create(user *model.User) error                 // Saves a new user to database
	FindByEmail(email string) (*model.User, error) // Finds user by email address
}

// userRepository is the actual implementation that does the real database work
// This is like the kitchen that prepares the food from the menu
type userRepository struct {
	db *gorm.DB // Database connection - our "kitchen tools"
}

// NewUserRepository is like a factory that creates our ready-to-use database operations helper
// We give it a database connection (db) and it gives us back a UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	// Create new userRepository instance with the provided database connection
	// We return it as UserRepository interface type (like putting our kitchen behind a service window)
	return &userRepository{db}
}

// Create saves a new user to the database
func (r *userRepository) Create(user *model.User) error {
	// r.db.Create performs the actual SQL INSERT operation
	// .Error gives us any error that might have occurred
	return r.db.Create(user).Error
}

// FindByEmail looks up a user by their email address
func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User // Empty container to store the found user

	// This is like doing: SELECT * FROM users WHERE email = ? LIMIT 1
	err := r.db.Where("email = ?", email).First(&user).Error

	// Return pointer to the found user and any error (nil if no error)
	return &user, err
}
