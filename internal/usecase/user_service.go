package usecase

import (
	"github.com/wahyunurdian26/cst_app_new/internal/entity"
)

type UserService interface {
	Create(user *entity.User) error
	GetById(id uint) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id uint) error
	GetAll() ([]entity.User, error)
}
