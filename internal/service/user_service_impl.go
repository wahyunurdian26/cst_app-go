package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"github.com/wahyunurdian26/cst_app_new/internal/model"
	"github.com/wahyunurdian26/cst_app_new/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	UserRepository repository.UserRepository
	Log            *logrus.Logger
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validator *validator.Validate, log *logrus.Logger) UserService {
	return &userService{
		UserRepository: userRepository,
		Validate:       validator,
		Log:            log,
	}
}

func (u *userService) Create(ctx context.Context, request *model.UserCreateRequest) (*entity.User, error) {
	// Validate request
	if err := u.Validate.Struct(request); err != nil {
		u.Log.Warnf("User creation failed: validation error: %v", err)
		return nil, fiber.ErrBadRequest
	}

	// Check if user already exists
	existingUser, err := u.UserRepository.FindByEmailOrUsername(request.Email, request.Username)
	if err != nil {
		u.Log.Warnf("User creation failed: database error: %v", err)
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}

	if existingUser != nil {
		u.Log.Warnf("User creation failed: user already exists with email %s or username %s", request.Email, request.Username)
		return nil, fiber.ErrConflict
	}

	// Hash password
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		u.Log.Warnf("User creation failed: password hashing error: %v", err)
		return nil, fiber.ErrInternalServerError
	}

	// Create new user entity
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

	// Save user to database
	if err := u.UserRepository.Create(user); err != nil {
		u.Log.Warnf("User creation failed: database error: %v", err)
		return nil, fiber.ErrInternalServerError
	}

	u.Log.Infof("User created successfully: ID %d, Email: %s, Username: %s", user.ID, user.Email, user.Username)

	// Kembalikan entity.User sesuai dengan deklarasi interface
	return user, nil
}

func (u *userService) GetById(id uuid.UUID) (*entity.User, error) {
	user, err := u.UserRepository.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.ErrNotFound
		}
		u.Log.Warnf("Failed to get user by ID %d: %v", id, err)
		return nil, err
	}
	return user, nil
}

func (u *userService) Update(ctx context.Context, request *model.UserUpdateRequest) (*entity.User, error) {
	if request == nil {
		u.Log.Warn("Update failed: missing request body")
		return nil, fiber.NewError(fiber.StatusBadRequest, "Request body is required")
	}

	if err := u.Validate.Struct(request); err != nil {
		u.Log.Warnf("Update failed: validation error: %v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid request data")
	}

	user, err := u.UserRepository.GetById(request.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.Log.Warnf("Update failed: user with ID %d not found", request.Id)
			return nil, fiber.ErrNotFound
		}
		u.Log.Warnf("Update failed: database error for user ID %d: %v", request.Id, err)
		return nil, err
	}

	if user == nil {
		u.Log.Warnf("Update failed: user with ID %d not found", request.Id)
		return nil, fiber.ErrNotFound
	}

	// Update fields if provided
	if request.Username != "" {
		user.Username = request.Username
	}
	if request.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			u.Log.Warn("Update failed: password hashing error")
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
	user.StatusActive = request.StatusActive
	if request.IDBusinessGroupDigital != "" {
		user.IDBusinessGroupDigital = request.IDBusinessGroupDigital
	}

	// Save updated user to database
	if err := u.UserRepository.Update(user); err != nil {
		u.Log.Warnf("Update failed: database error for user ID %d: %v", request.Id, err)
		return nil, fiber.ErrInternalServerError
	}

	u.Log.Infof("User updated successfully: ID %d", request.Id)

	return user, nil // ⬅️ Hanya return entity.User, tanpa response tambahan
}

func (u *userService) Delete(id uuid.UUID) error {
	existingUser, err := u.UserRepository.GetById(id)
	if err != nil || existingUser == nil {
		u.Log.Warnf("Delete failed: user ID %d not found", id)
		return errors.New("user not found")
	}
	return u.UserRepository.Delete(id)
}

func (u *userService) GetAll() ([]entity.User, error) {
	users, err := u.UserRepository.GetAll()
	if err != nil {
		u.Log.Warnf("Failed to retrieve users: %v", err)
		return nil, err
	}
	return users, nil
}
