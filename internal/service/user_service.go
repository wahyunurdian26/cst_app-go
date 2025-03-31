package service

import (
	"context"

	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"github.com/wahyunurdian26/cst_app_new/internal/model"
)

type UserService interface {
	Create(ctx context.Context, request *model.UserCreateRequest) (*entity.User, error) // Sesuaikan return type
	GetById(id uint) (*entity.User, error)
	Update(ctx context.Context, request *model.UserUpdateRequest) (*entity.User, error)
	Delete(id uint) error
	GetAll() ([]entity.User, error)
}
