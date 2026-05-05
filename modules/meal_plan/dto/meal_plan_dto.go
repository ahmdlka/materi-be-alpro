package dto

import (
	"github.com/google/uuid"
	"github.com/Mobilizes/materi-be-alpro/database/entities"
)

type GenerateMealPlanRequest struct {
	Mode             string            `json:"mode" binding:"required,oneof=FAST SPECIFIC"`
	DurationDays     int               `json:"duration_days" binding:"required,min=1,max=30"`
	ExtraConstraints *ExtraConstraints `json:"extra_constraints"`
}

type ExtraConstraints struct {
	BudgetPerDay       int      `json:"budget_per_day"`
	ExcludeIngredients []string `json:"exclude_ingredients"`
	PreferLocalFood    bool     `json:"prefer_local_food"`
}

type MealPlanResponse struct {
	ID           uuid.UUID             `json:"meal_plan_id"`
	Mode         string                `json:"mode"`
	Version      int                   `json:"version"`
	DurationDays int                   `json:"duration_days"`
	Plan         entities.FullPlanData `json:"plan"`
	IsActive     bool                  `json:"is_active"`
}

type MealPlanVersionResponse struct {
	ID         uuid.UUID             `json:"id"`
	Version    int                   `json:"version"`
	PlanData   entities.FullPlanData `json:"plan"`
	ChangeNote string                `json:"change_note"`
	CreatedAt  string                `json:"created_at"`
}
