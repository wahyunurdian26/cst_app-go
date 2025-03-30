package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"github.com/wahyunurdian26/cst_app_new/internal/model"
	"github.com/wahyunurdian26/cst_app_new/internal/repository"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	UserRepository repository.UserRepository

	Validate *validator.Validate
}

// NewUserService mengembalikan instance dari UserService
func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		UserRepository: userRepository,
		Validate:       validator.New(),
	}
}

// Create membuat user baru dengan validasi dan hashing password
func (c *userService) Create(ctx context.Context, request *model.UserCreateRequest) (*model.UserResponse, error) {
	// Validasi request
	err := c.Validate.Struct(request)
	fmt.Println("invalid request body", err)
	if err != nil {
		return nil, fiber.ErrBadRequest
	}

	//cek apakah sudah ada
	existingUser, err := c.UserRepository.FindByEmailOrUsername(request.Email, request.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}

	if existingUser != nil { // Jika user sudah ada
		return nil, fiber.ErrConflict
	}

	// Hash password
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
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
	if err := c.UserRepository.Create(user); err != nil {

		return nil, fiber.ErrInternalServerError
	}

	// Response sukses
	return &model.UserResponse{
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
func (s *userService) GetById(id uint) (*entity.User, error) {
	user, err := s.UserRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fiber.ErrNotFound 
	}
	return user, nil
}

// Implementasi metode Update dengan validasi
func (c *userService) Update(ctx context.Context, request *model.UserUpdateRequest) (*model.UserResponse, error) {
	// Validasi request
	if err := c.Validate.Struct(request); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid request data")
	}

	// Cek apakah user ada di database
	user, err := c.UserRepository.GetById(request.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fiber.ErrNotFound
		}
		return nil, err // Error lain
	}

	// Update hanya field yang diisi (Gunakan GORM `Updates`)
	updateData := map[string]interface{}{}
	if request.Username != "" {
		updateData["username"] = request.Username
	}
	if request.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, fiber.ErrInternalServerError
		}
		user.Password = string(password)

	}
	if request.IDRole != "" {
		updateData["id_role"] = request.IDRole
	}
	if request.IDBusinessGroup != "" {
		updateData["id_business_group"] = request.IDBusinessGroup
	}
	if request.IDSubBusinessGroup != "" {
		updateData["id_sub_business_group"] = request.IDSubBusinessGroup
	}
	if request.EmailPIC != "" {
		updateData["email_pic"] = request.EmailPIC
	}
	updateData["status_active"] = request.StatusActive // Tetap di-update walaupun default-nya false
	if request.IDBusinessGroupDigital != "" {
		updateData["id_business_group_digital"] = request.IDBusinessGroupDigital
	}

	// Jalankan update di database

	if err := c.UserRepository.Update(user); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return &model.UserResponse{
		Code:    fiber.StatusOK,
		Status:  "UPDATED",
		Message: "User successfully updated",
		Data:    updateData,
	}, nil
}

// Implementasi metode Delete dengan validasi
func (s *userService) Delete(id uint) error {
	// Pastikan user ada sebelum dihapus
	existingUser, err := s.UserRepository.GetById(id)
	if err != nil || existingUser == nil {
		return errors.New("user not found")
	}

	return s.UserRepository.Delete(id)
}

// Implementasi metode GetAll untuk mendapatkan semua user
func (s *userService) GetAll() ([]entity.User, error) {
	users, err := s.UserRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
