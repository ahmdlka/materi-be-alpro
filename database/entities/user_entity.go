// database/entities/user_entity.go

package entities

import (
	"github.com/lib/pq"
)

type User struct {
	Common
	Email     string      `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password  string      `gorm:"type:varchar(255);not null"             json:"-"`
	Name      string      `gorm:"type:varchar(255)"                      json:"name"`
	Profile   UserProfile `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"profile,omitempty"`
	MealPlans []MealPlan  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"meal_plans,omitempty"`
}

type UserProfile struct {
	Common
	UserID          string         `gorm:"type:uuid;not null;index"         json:"user_id"`
	Age             int            `gorm:"type:int"                         json:"age"`
	WeightKg        float64        `gorm:"type:float"                       json:"weight_kg"`
	HeightCm        float64        `gorm:"type:float"                       json:"height_cm"`
	Gender          string         `gorm:"type:varchar(10)"                 json:"gender"`
	ActivityLevel   string         `gorm:"type:varchar(50)"                 json:"activity_level"` // sedentary, light, moderate, active
	Goal            string         `gorm:"type:varchar(50)"                 json:"goal"`           // weight_loss, muscle_gain, maintain
	Allergies       pq.StringArray `gorm:"type:text[]" json:"allergies"`
	Diseases        pq.StringArray `gorm:"type:text[]" json:"diseases"`
	FoodPreferences pq.StringArray `gorm:"type:text[]" json:"food_preferences"`

	BudgetPerDay    int            `gorm:"type:int"                         json:"budget_per_day"` // dalam Rupiah
}
