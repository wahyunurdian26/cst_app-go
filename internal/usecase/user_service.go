package usecase

import (
	"context"

	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"github.com/wahyunurdian26/cst_app_new/internal/model"
)

type UserService interface {
	Create(ctx context.Context, request *model.UserCreateRequest) (*model.Response, error)
	GetById(id uint) (*entity.User, error)
	Update(ctx context.Context, req *model.UserUpdateRequest) (*model.Response, error)
	Delete(id uint) error
	GetAll() ([]entity.User, error)
}
