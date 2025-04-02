package repository

import (
	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"gorm.io/gorm"
)

type authRepositoryiImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepositoryiImpl{DB: db}
}

func (r authRepositoryiImpl) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
