package repository

import (
	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}
