package repository

import (
	"errors"
	"github.com/shahid-io/shippix/pkg/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetByEmailOrPhone(emailOrPhone string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ? OR phone = ?", emailOrPhone, emailOrPhone).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}
