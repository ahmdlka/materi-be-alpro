package repository

import (
	"github.com/Mobilizes/materi-be-alpro/database/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) Create(history *entities.ChatHistory) error {
	return r.db.Create(history).Error
}

func (r *ChatRepository) FindByUserID(userID uuid.UUID) ([]entities.ChatHistory, error) {
	var history []entities.ChatHistory
	err := r.db.Where("user_id = ?", userID).Order("created_at asc").Find(&history).Error
	return history, err
}
