package usecase

import "github.com/wahyunurdian26/cst_app_new/internal/repository"

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}
