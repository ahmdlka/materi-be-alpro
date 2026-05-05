// database/entities/chat_history_entity.go

package entities

import "github.com/google/uuid"

type ChatHistory struct {
    Common
    UserID     uuid.UUID  `gorm:"type:uuid;not null;index"                              json:"user_id"`
    MealPlanID *uuid.UUID `gorm:"type:uuid;index"                                       json:"meal_plan_id"` // nullable
    Role       string     `gorm:"type:varchar(10)"                                      json:"role"`         // user | assistant
    Message    string     `gorm:"type:text;not null"                                    json:"message"`
    Intent     string     `gorm:"type:varchar(20)"                                      json:"intent"`       // REFINE_MENU | ASK_QUESTION
}
