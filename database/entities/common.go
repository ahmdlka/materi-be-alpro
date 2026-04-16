package entities

import (
    "time"
    "gorm.io/gorm"
)

// Common model equivalent of gorm.Model
type Common struct {
    ID        uint           `gorm:"primarykey" json:"id"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
