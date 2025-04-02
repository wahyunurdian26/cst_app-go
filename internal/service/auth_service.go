package service

import (
	"github.com/wahyunurdian26/cst_app_new/internal/model"
)

type AuthService interface {
	Login(req model.LoginUserRequest) (string, error)
}
