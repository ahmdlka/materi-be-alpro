package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Mobilizes/materi-be-alpro/database/entities"
	"github.com/Mobilizes/materi-be-alpro/modules/meal_plan/dto"
	"github.com/Mobilizes/materi-be-alpro/modules/meal_plan/repository"
	userRepo "github.com/Mobilizes/materi-be-alpro/modules/user/repository"
	"github.com/Mobilizes/materi-be-alpro/pkg/ragclient"
	"github.com/google/uuid"
)

type MealPlanService struct {
	repo      *repository.MealPlanRepository
	userRepo  *userRepo.UserRepository
	ragClient *ragclient.RAGClient
}

func NewMealPlanService(repo *repository.MealPlanRepository, userRepo *userRepo.UserRepository, ragClient *ragclient.RAGClient) *MealPlanService {
	return &MealPlanService{
		repo:      repo,
		userRepo:  userRepo,
		ragClient: ragClient,
	}
}

func (s *MealPlanService) Generate(userID uuid.UUID, req *dto.GenerateMealPlanRequest) (*dto.MealPlanResponse, error) {
	// 1. Get user profile
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	if user.Profile.ID == uuid.Nil {
		return nil, errors.New("user profile is incomplete, please update your profile first")
	}

	// 2. Prepare payload for RAG
	payload := map[string]interface{}{
		"user_profile": map[string]interface{}{
			"age":              user.Profile.Age,
			"weight_kg":        user.Profile.WeightKg,
			"height_cm":        user.Profile.HeightCm,
			"gender":           user.Profile.Gender,
			"activity_level":   user.Profile.ActivityLevel,
			"goal":             user.Profile.Goal,
			"allergies":        user.Profile.Allergies,
			"diseases":         user.Profile.Diseases,
			"food_preferences": user.Profile.FoodPreferences,
		},
		"constraints": map[string]interface{}{
			"duration_days": req.DurationDays,
		},
	}

	if req.ExtraConstraints != nil {
		constraints := payload["constraints"].(map[string]interface{})
		constraints["budget_per_day"] = req.ExtraConstraints.BudgetPerDay
		constraints["exclude_ingredients"] = req.ExtraConstraints.ExcludeIngredients
		constraints["prefer_local_food"] = req.ExtraConstraints.PreferLocalFood
	}

	// 3. Call RAG Service
	ragRes, err := s.ragClient.Generate(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to generate meal plan from AI: %w", err)
	}

	planData, ok := ragRes["data"].(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid response from AI service")
	}

	// 4. Save to Database
	mealPlan := &entities.MealPlan{
		UserID:       userID.String(),
		Mode:         req.Mode,
		DurationDays: req.DurationDays,
		PlanData:     planData,
		Version:      1,
		IsActive:     true,
	}

	if err := s.repo.Create(mealPlan); err != nil {
		return nil, fmt.Errorf("failed to save meal plan: %w", err)
	}

	// 5. Save Version
	version := &entities.MealPlanVersion{
		MealPlanID: mealPlan.ID.String(),
		Version:    1,
		PlanData:   planData,
		ChangeNote: "Initial generation",
	}
	s.repo.CreateVersion(version)

	return s.mapToMealPlanResponse(mealPlan), nil
}

func (s *MealPlanService) GetByUserID(userID uuid.UUID) ([]dto.MealPlanResponse, error) {
	plans, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var res []dto.MealPlanResponse
	for _, p := range plans {
		res = append(res, *s.mapToMealPlanResponse(&p))
	}
	return res, nil
}

func (s *MealPlanService) GetByID(id uuid.UUID) (*dto.MealPlanResponse, error) {
	plan, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return s.mapToMealPlanResponse(plan), nil
}

func (s *MealPlanService) GetVersions(mealPlanID uuid.UUID) ([]dto.MealPlanVersionResponse, error) {
	versions, err := s.repo.FindVersionsByMealPlanID(mealPlanID)
	if err != nil {
		return nil, err
	}

	var res []dto.MealPlanVersionResponse
	for _, v := range versions {
		res = append(res, dto.MealPlanVersionResponse{
			ID:         v.ID,
			Version:    v.Version,
			PlanData:   s.convertToFullPlanData(v.PlanData),
			ChangeNote: v.ChangeNote,
			CreatedAt:  v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return res, nil
}

func (s *MealPlanService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *MealPlanService) mapToMealPlanResponse(p *entities.MealPlan) *dto.MealPlanResponse {
	return &dto.MealPlanResponse{
		ID:           p.ID,
		Mode:         p.Mode,
		Version:      p.Version,
		DurationDays: p.DurationDays,
		Plan:         s.convertToFullPlanData(p.PlanData),
		IsActive:     p.IsActive,
	}
}

func (s *MealPlanService) convertToFullPlanData(data entities.PlanData) entities.FullPlanData {
	var fullPlan entities.FullPlanData
	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &fullPlan)
	return fullPlan
}
