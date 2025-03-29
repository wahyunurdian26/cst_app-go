package repository

import (
	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"gorm.io/gorm"
)

// Implementasi UserRepository
type userRepository struct {
	db *gorm.DB
}

// Konstruktor untuk membuat UserRepository baru
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Implementasi metode Create
func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

// Implementasi metode GetById
func (r *userRepository) GetById(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	return &user, err
}

// Implementasi metode Update
func (r *userRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

// Implementasi metode Delete
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}

// Implementasi metode GetAll
func (r *userRepository) GetAll() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	return users, err
}
