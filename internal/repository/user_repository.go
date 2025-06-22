package repository

import (
	"github.com/dassajib/prohor-api/internal/model"
	"gorm.io/gorm"
)

// this interface is set what db related operations we support
type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

// this struct is the actual implementation that does the real database work
type userRepository struct {
	db *gorm.DB
}

// repository constructor function
func NewUserRepository(db *gorm.DB) UserRepository {
	// create a new UserRepository instance with provided db connection
	return &userRepository{db}
}

// receiver func to create user
func (r *userRepository) Create(user *model.User) error {
	// sql insert operation's value and error that might occurred
	return r.db.Create(user).Error
}

// receiver func to find the user by given email
func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	// sql operation to find the user exist or not
	err := r.db.Where("email = ?", email).First(&user).Error
	// return the pointer to the user orr error that might occurred
	return &user, err
}
