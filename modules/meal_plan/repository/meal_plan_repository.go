package repository

import (
	"github.com/Mobilizes/materi-be-alpro/database/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MealPlanRepository struct {
	db *gorm.DB
}

func NewMealPlanRepository(db *gorm.DB) *MealPlanRepository {
	return &MealPlanRepository{db: db}
}

func (r *MealPlanRepository) Create(plan *entities.MealPlan) error {
	return r.db.Create(plan).Error
}

func (r *MealPlanRepository) CreateVersion(version *entities.MealPlanVersion) error {
	return r.db.Create(version).Error
}

func (r *MealPlanRepository) FindByID(id uuid.UUID) (*entities.MealPlan, error) {
	var plan entities.MealPlan
	err := r.db.Where("id = ?", id).First(&plan).Error
	return &plan, err
}

func (r *MealPlanRepository) FindByUserID(userID uuid.UUID) ([]entities.MealPlan, error) {
	var plans []entities.MealPlan
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&plans).Error
	return plans, err
}

func (r *MealPlanRepository) FindVersionsByMealPlanID(mealPlanID uuid.UUID) ([]entities.MealPlanVersion, error) {
	var versions []entities.MealPlanVersion
	err := r.db.Where("meal_plan_id = ?", mealPlanID).Order("version desc").Find(&versions).Error
	return versions, err
}

func (r *MealPlanRepository) Update(plan *entities.MealPlan) error {
	return r.db.Save(plan).Error
}

func (r *MealPlanRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.MealPlan{}, id).Error
}
