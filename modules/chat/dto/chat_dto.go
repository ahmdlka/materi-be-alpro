package dto

import (
	"github.com/google/uuid"
	"github.com/Mobilizes/materi-be-alpro/database/entities"
)

type ChatRequest struct {
	Message    string     `json:"message" binding:"required"`
	MealPlanID *uuid.UUID `json:"meal_plan_id"`
}

type ChatResponse struct {
	Reply           string                `json:"reply"`
	Intent          string                `json:"intent"`
	UpdatedMealPlan entities.FullPlanData `json:"updated_meal_plan,omitempty"`
	NewVersion      int                   `json:"new_version,omitempty"`
}

type ChatHistoryResponse struct {
	ID         uuid.UUID  `json:"id"`
	MealPlanID *uuid.UUID `json:"meal_plan_id"`
	Role       string     `json:"role"`
	Message    string     `json:"message"`
	Intent     string     `json:"intent"`
	CreatedAt  string     `json:"created_at"`
}
