package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"github.com/wahyunurdian26/cst_app_new/internal/model"
)

type UserService interface {
	Create(ctx context.Context, request *model.UserCreateRequest) (*entity.User, error)
	GetById(id uuid.UUID) (*entity.User, error)
	Update(ctx context.Context, request *model.UserUpdateRequest) (*entity.User, error)
	Delete(id uuid.UUID) error
	GetAll() ([]entity.User, error)
}
