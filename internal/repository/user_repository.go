package repository

import (
	"github.com/google/uuid"
	"github.com/wahyunurdian26/cst_app_new/internal/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmailOrUsername(email, username string) (*entity.User, error)
	GetById(id uuid.UUID) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id uuid.UUID) error
	GetAll() ([]entity.User, error)
}


	



	