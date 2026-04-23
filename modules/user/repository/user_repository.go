package repository

import (
	"github.com/Mobilizes/materi-be-alpro/database/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindByID(id int) (*entities.User, error) {
	var user entities.User
	err := r.db.Select("id", "name", "email", "role").Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindAll() ([]entities.User, error) {
	var users []entities.User
	result := r.db.Select("id", "name", "email", "role").Find(&users)
	return users, result.Error
}
