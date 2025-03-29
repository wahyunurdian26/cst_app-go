package usecase

import (
	"errors"
	"strings"

	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"github.com/wahyunurdian26/cst_app_new/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repository.UserRepository
}

// NewUserService mengembalikan instance dari UserService
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// Implementasi metode Create dengan validasi dan hashing password
func (s *userService) Create(user *entity.User) error {
	// Validasi input
	if user.Email == "" || !strings.Contains(user.Email, "@") {
		return errors.New("invalid email format")
	}
	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	// Hash password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Simpan ke repository
	return s.repo.Create(user)
}

// Implementasi metode GetById dengan pengecekan apakah user ada
func (s *userService) GetById(id uint) (*entity.User, error) {
	user, err := s.repo.GetById(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// Implementasi metode Update dengan validasi
func (s *userService) Update(userID uint, updates map[string]interface{}) error {
	// Pastikan user ada sebelum diupdate
	existingUser, err := s.repo.GetById(userID)
	if err != nil || existingUser == nil {
		return errors.New("user not found")
	}

	// Jika password ingin diupdate, hash dulu
	if password, ok := updates["password"].(string); ok && password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		updates["password"] = string(hashedPassword)
	} else {
		delete(updates, "password") // Jangan update password jika tidak dikirim
	}

	// Lakukan update dengan GORM
	return s.repo.UpdateFields(userID, updates)
}

// Implementasi metode Delete dengan validasi
func (s *userService) Delete(id uint) error {
	// Pastikan user ada sebelum dihapus
	existingUser, err := s.repo.GetById(id)
	if err != nil || existingUser == nil {
		return errors.New("user not found")
	}

	return s.repo.Delete(id)
}

// Implementasi metode GetAll untuk mendapatkan semua user
func (s *userService) GetAll() ([]entity.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
