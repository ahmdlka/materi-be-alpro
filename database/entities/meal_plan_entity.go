// database/entities/meal_plan_entity.go

package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// --- Helper Structs for PlanData (Menu Book Style) ---

type Ingredient struct {
	Name          string  `json:"name"`
	Weight        float64 `json:"weight"`
	Unit          string  `json:"unit"`
	Calories      float64 `json:"calories"`
	Protein       float64 `json:"protein"`
	Carbs         float64 `json:"carbs"`
	Fat           float64 `json:"fat"`
	PriceEstimate int     `json:"price_estimate"`
}

type MealNutrition struct {
	Calories float64 `json:"calories"`
	Protein  float64 `json:"protein"`
	Carbs    float64 `json:"carbs"`
	Fat      float64 `json:"fat"`
}

type MealDetail struct {
	Type             string        `json:"type"` // breakfast, lunch, dinner, snack
	Name             string        `json:"name"`
	BudgetEstimate   int           `json:"budget_estimate"`
	NutritionSummary MealNutrition `json:"nutrition_summary"`
	Ingredients      []Ingredient  `json:"ingredients"`
	Instructions     []string      `json:"instructions"`
}

type DayPlan struct {
	Day              int           `json:"day"`
	Meals            []MealDetail  `json:"meals"`
	DailyTotalBudget int           `json:"daily_total_budget"`
	DailyNutrition   MealNutrition `json:"daily_nutrition"`
}

type PlanSummary struct {
	AvgDailyCalories     float64 `json:"avg_daily_calories"`
	AvgProtein           float64 `json:"avg_protein"`
	AvgCarbs             float64 `json:"avg_carbs"`
	AvgFat               float64 `json:"avg_fat"`
	TotalEstimatedBudget int     `json:"total_estimated_budget"`
}

// FullPlanData adalah struktur lengkap dari PlanData
type FullPlanData struct {
	Days             []DayPlan   `json:"days"`
	NutritionSummary PlanSummary `json:"nutrition_summary"`
}

// ---

// PlanData adalah custom type untuk menyimpan JSONB di PostgreSQL via GORM
type PlanData map[string]interface{}

func (p PlanData) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *PlanData) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan PlanData")
	}
	return json.Unmarshal(bytes, p)
}

// ---

type MealPlan struct {
	Common
	UserID       string            `gorm:"type:uuid;not null;index"                              json:"user_id"`
	Mode         string            `gorm:"type:varchar(10)"                                      json:"mode"` // FAST | SPECIFIC
	DurationDays int               `gorm:"default:7"                                             json:"duration_days"`
	PlanData     PlanData          `gorm:"type:jsonb;not null"                                   json:"plan"`
	Version      int               `gorm:"default:1"                                             json:"version"`
	IsActive     bool              `gorm:"default:true"                                          json:"is_active"`
	Versions     []MealPlanVersion `gorm:"foreignKey:MealPlanID;constraint:OnDelete:CASCADE"   json:"versions,omitempty"`
}

type MealPlanVersion struct {
	Common
	MealPlanID string   `gorm:"type:uuid;not null;index"  json:"meal_plan_id"`
	Version    int      `gorm:"type:int"                  json:"version"`
	PlanData   PlanData `gorm:"type:jsonb;not null"       json:"plan"`
	ChangeNote string   `gorm:"type:text"                 json:"change_note"` // "ganti ayam jadi tempe"
}
