package repository

import (
	"github.com/Mobilizes/materi-be-alpro/database/entities"
	"github.com/google/uuid"
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

func (r *UserRepository) FindByID(id uuid.UUID) (*entities.User, error) {
	var user entities.User
	err := r.db.Preload("Profile").Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindAll() ([]entities.User, error) {
	var users []entities.User
	result := r.db.Preload("Profile").Find(&users)
	return users, result.Error
}

func (r *UserRepository) UpdateProfile(profile *entities.UserProfile) error {
	return r.db.Save(profile).Error
}

func (r *UserRepository) GetProfileByUserID(userID uuid.UUID) (*entities.UserProfile, error) {
	var profile entities.UserProfile
	err := r.db.Where("user_id = ?", userID).First(&profile).Error
	return &profile, err
}
