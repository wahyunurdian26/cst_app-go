package repository

import (
	"github.com/wahyunurdian26/cst_app_new/internal/entity"
)

type AuthRepository interface {
	FindByEmail(email string) (*entity.User, error)
}
