package repository

import (
	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"gorm.io/gorm"
)

// Implementasi dengan nama berbeda
type UserRepositoryImpl struct {
	DB *gorm.DB
}

// Konstruktor
func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db} // ✅ Benar, menggunakan pointer
}

// Implementasi metode Create
func (r *UserRepositoryImpl) Create(user *entity.User) error {
	return r.DB.Create(user).Error
}

// Implementasi metode FindByEmailOrUsername
func (r *UserRepositoryImpl) FindByEmailOrUsername(email, username string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("email = ? OR username = ?", email, username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Implementasi metode GetById
func (r *UserRepositoryImpl) GetById(id uint) (*entity.User, error) {
	var user entity.User
	err := r.DB.Where("id = ?", id).Take(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // ✅ Konsisten dengan FindByEmailOrUsername
		}
		return nil, err
	}
	return &user, nil
}

// Implementasi metode Update
func (r *UserRepositoryImpl) Update(user *entity.User) error {
	return r.DB.Save(user).Error
}

// Implementasi metode Delete
func (r *UserRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&entity.User{}, id).Error
}

// Implementasi metode GetAll
func (r *UserRepositoryImpl) GetAll() ([]entity.User, error) {
	var users []entity.User
	err := r.DB.Find(&users).Error
	return users, err
}
