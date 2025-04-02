package service

import (
	"errors"

	"github.com/wahyunurdian26/cst_app_new/internal/model"
	"github.com/wahyunurdian26/cst_app_new/internal/repository"
	"github.com/wahyunurdian26/cst_app_new/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type authServiceImpl struct {
	authRepo repository.AuthRepository
}

func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return &authServiceImpl{authRepo: authRepo}
}

func (r *authServiceImpl) Login(req model.LoginUserRequest) (string, error) {
	user, err := r.authRepo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, user.IDRole, user.IDBusinessGroup, user.IDSubBusinessGroup, user.IDBusinessGroupDigital)
	if err != nil {
		return "", err
	}

	return token, nil

}
