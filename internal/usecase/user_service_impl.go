package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"github.com/wahyunurdian26/cst_app_new/internal/model"
	"github.com/wahyunurdian26/cst_app_new/internal/repository"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	UserRepository repository.UserRepository
	Log            *logrus.Logger
	Validate       *validator.Validate
}

// NewUserService mengembalikan instance dari UserService
func NewUserService(userRepository repository.UserRepository, validator *validator.Validate, log *logrus.Logger) UserService {
	return &userService{
		UserRepository: userRepository,
		Validate:       validator,
		Log:            log,
	}
}

// Create membuat user baru dengan validasi dan hashing password
func (u *userService) Create(ctx context.Context, request *model.UserCreateRequest) (*model.Response, error) {
	// Validasi request
	err := u.Validate.Struct(request)

	if err != nil {
		u.Log.Warn("Failed to validate  request body")
		return nil, fiber.ErrBadRequest

	}

	//cek apakah sudah ada
	existingUser, err := u.UserRepository.FindByEmailOrUsername(request.Email, request.Username)
	if err != nil {
		u.Log.Warn("Failed to FindUser")
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}

	if existingUser != nil { // Jika user sudah ada
		return nil, fiber.ErrConflict
	}

	// Hash password
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		u.Log.Warn("Failed to generate password")
		return nil, fiber.ErrInternalServerError
	}

	// Buat user
	user := &entity.User{
		Email:                  request.Email,
		Username:               request.Username,
		Password:               string(password),
		IDRole:                 request.IDRole,
		IDBusinessGroup:        request.IDBusinessGroup,
		IDSubBusinessGroup:     request.IDSubBusinessGroup,
		EmailPIC:               request.EmailPIC,
		StatusActive:           request.StatusActive,
		IDBusinessGroupDigital: request.IDBusinessGroupDigital,
	}

	// Simpan ke database
	if err := u.UserRepository.Create(user); err != nil {
		u.Log.Warn("Failed to create userr")
		return nil, fiber.ErrInternalServerError
	}

	// Response sukses
	return &model.Response{
		Code:    fiber.StatusCreated,
		Status:  "CREATED",
		Message: "User successfully created",
		Data: map[string]interface{}{
			"id":                        user.ID,
			"email":                     user.Email,
			"username":                  user.Username,
			"id_role":                   user.IDRole,
			"id_business_group":         user.IDBusinessGroup,
			"id_sub_business_group":     user.IDSubBusinessGroup,
			"email_pic":                 user.EmailPIC,
			"status_active":             user.StatusActive,
			"id_business_group_digital": user.IDBusinessGroupDigital,
		},
	}, nil
}

// Implementasi metode GetById dengan pengecekan apakah user ada
func (u *userService) GetById(id uint) (*entity.User, error) {
	user, err := u.UserRepository.GetById(id)
	if err != nil {
		u.Log.Warn("Failed to GetById")
		return nil, err
	}
	return user, nil
}

// Implementasi metode Update dengan validasi
func (u *userService) Update(ctx context.Context, request *model.UserUpdateRequest) (*model.Response, error) {
	// Validasi request
	if err := u.Validate.Struct(request); err != nil {
		u.Log.Warn("Failed to validate user request")
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid request data")
	}

	// Cek apakah user ada di database

	user, err := u.UserRepository.GetById(request.Id)
	if err != nil {
		u.Log.Warn("Failed to GetById")
		if err == gorm.ErrRecordNotFound {
			return nil, fiber.ErrNotFound
		}
		return nil, err // Error lain
	}

	// Update hanya field yang diisi (Gunakan GORM `Updates`)

	if request.Username != "" {

		user.Username = request.Username
	}
	if request.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			u.Log.Warn("Failed to generate hash password")
			return nil, fiber.ErrInternalServerError
		}
		user.Password = string(password)

	}
	if request.IDRole != "" {
		user.IDRole = request.IDRole
	}
	if request.IDBusinessGroup != "" {
		user.IDBusinessGroup = request.IDBusinessGroup
	}
	if request.IDSubBusinessGroup != "" {
		user.IDSubBusinessGroup = request.IDSubBusinessGroup
	}
	if request.EmailPIC != "" {
		user.EmailPIC = request.EmailPIC
	}
	user.StatusActive = request.StatusActive // Tetap di-update walaupun default-nya false
	if request.IDBusinessGroupDigital != "" {
		user.IDBusinessGroupDigital = request.IDBusinessGroupDigital
	}

	// Jalankan update di database

	if err := u.UserRepository.Update(user); err != nil {
		u.Log.Warn("Failed to save user")
		return nil, fiber.ErrInternalServerError
	}

	return &model.Response{
		Code:    fiber.StatusOK,
		Status:  "UPDATED",
		Message: "User successfully updated",
		Data:    user,
	}, nil
}

// Implementasi metode Delete dengan validasi
func (u *userService) Delete(id uint) error {
	// Pastikan user ada sebelum dihapus
	existingUser, err := u.UserRepository.GetById(id)
	if err != nil || existingUser == nil {
		u.Log.Warn("Failed to find user")
		return errors.New("user not found")
	}

	return u.UserRepository.Delete(id)
}

// Implementasi metode GetAll untuk mendapatkan semua user
func (u *userService) GetAll() ([]entity.User, error) {
	users, err := u.UserRepository.GetAll()
	if err != nil {
		u.Log.Warn("Failed to get all user")
		return nil, err
	}
	return users, nil
}
