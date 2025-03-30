package repository

import "github.com/wahyunurdian26/cst_app_new/internal/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmailOrUsername(email, username string) (*entity.User, error)
	GetById(id uint) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id uint) error
	GetAll() ([]entity.User, error)
}
