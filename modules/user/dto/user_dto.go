package dto

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type CreateUserRequest struct {
	Name     string `json:"name"     binding:"required"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UpdateProfileRequest struct {
	Age             int            `json:"age"               binding:"required"`
	WeightKg        float64        `json:"weight_kg"         binding:"required"`
	HeightCm        float64        `json:"height_cm"         binding:"required"`
	Gender          string         `json:"gender"            binding:"required"`
	ActivityLevel   string         `json:"activity_level"    binding:"required"`
	Goal            string         `json:"goal"              binding:"required"`
	Allergies       pq.StringArray `json:"allergies"`
	Diseases        pq.StringArray `json:"diseases"`
	FoodPreferences pq.StringArray `json:"food_preferences"`
}

type UserResponse struct {
	ID      uuid.UUID        `json:"id"`
	Name    string           `json:"name"`
	Email   string           `json:"email"`
	Profile *ProfileResponse `json:"profile,omitempty"`
}

type ProfileResponse struct {
	Age             int            `json:"age"`
	WeightKg        float64        `json:"weight_kg"`
	HeightCm        float64        `json:"height_cm"`
	Gender          string         `json:"gender"`
	ActivityLevel   string         `json:"activity_level"`
	Goal            string         `json:"goal"`
	Allergies       pq.StringArray `json:"allergies"`
	Diseases        pq.StringArray `json:"diseases"`
	FoodPreferences pq.StringArray `json:"food_preferences"`
}
