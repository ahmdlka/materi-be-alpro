package service

import (
	"encoding/json"
	"fmt"

	"github.com/Mobilizes/materi-be-alpro/database/entities"
	"github.com/Mobilizes/materi-be-alpro/modules/chat/dto"
	"github.com/Mobilizes/materi-be-alpro/modules/chat/repository"
	mpRepo "github.com/Mobilizes/materi-be-alpro/modules/meal_plan/repository"
	userRepo "github.com/Mobilizes/materi-be-alpro/modules/user/repository"
	"github.com/Mobilizes/materi-be-alpro/pkg/ragclient"
	"github.com/google/uuid"
)

type ChatService struct {
	repo          *repository.ChatRepository
	intentService *IntentService
	mpRepo        *mpRepo.MealPlanRepository
	userRepo      *userRepo.UserRepository
	ragClient     *ragclient.RAGClient
}

func NewChatService(
	repo *repository.ChatRepository,
	intentService *IntentService,
	mpRepo *mpRepo.MealPlanRepository,
	userRepo *userRepo.UserRepository,
	ragClient *ragclient.RAGClient,
) *ChatService {
	return &ChatService{
		repo:          repo,
		intentService: intentService,
		mpRepo:        mpRepo,
		userRepo:      userRepo,
		ragClient:     ragClient,
	}
}

func (s *ChatService) ProcessMessage(userID uuid.UUID, req *dto.ChatRequest) (*dto.ChatResponse, error) {
	// 1. Identify Intent
	intent := s.intentService.Classify(req.MealPlanID)

	// 2. Get User Profile for Context
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	userProfile := map[string]interface{}{
		"age":       user.Profile.Age,
		"weight_kg": user.Profile.WeightKg,
		"height_cm": user.Profile.HeightCm,
		"gender":    user.Profile.Gender,
		"goal":      user.Profile.Goal,
	}

	// 3. Save User Message to History
	s.repo.Create(&entities.ChatHistory{
		UserID:     userID,
		MealPlanID: req.MealPlanID,
		Role:       "user",
		Message:    req.Message,
		Intent:     intent,
	})

	var response dto.ChatResponse
	response.Intent = intent

	if intent == IntentRefineMenu {
		// REFINE_MENU Logic
		plan, err := s.mpRepo.FindByID(*req.MealPlanID)
		if err != nil {
			return nil, fmt.Errorf("meal plan not found: %w", err)
		}

		ragPayload := map[string]interface{}{
			"meal_plan":    plan.PlanData,
			"instruction":  req.Message,
			"user_profile": userProfile,
		}

		ragRes, err := s.ragClient.Refine(ragPayload)
		if err != nil {
			return nil, err
		}

		updatedPlan := ragRes["data"].(map[string]interface{})
		newVersion := plan.Version + 1

		// Update MealPlan
		plan.PlanData = updatedPlan
		plan.Version = newVersion
		s.mpRepo.Update(plan)

		// Create Version History
		s.mpRepo.CreateVersion(&entities.MealPlanVersion{
			MealPlanID: plan.ID.String(),
			Version:    newVersion,
			PlanData:   updatedPlan,
			ChangeNote: req.Message,
		})

		response.Reply = ragRes["reply"].(string)
		if response.Reply == "" {
			response.Reply = "Saya telah memperbarui meal plan Anda sesuai instruksi."
		}
		
		// Convert to FullPlanData for DTO
		var fullPlan entities.FullPlanData
		bytes, _ := json.Marshal(updatedPlan)
		json.Unmarshal(bytes, &fullPlan)
		
		response.UpdatedMealPlan = fullPlan
		response.NewVersion = newVersion

	} else {
		// ASK_QUESTION Logic
		ragPayload := map[string]interface{}{
			"question":     req.Message,
			"user_profile": userProfile,
		}

		ragRes, err := s.ragClient.Ask(ragPayload)
		if err != nil {
			return nil, err
		}

		response.Reply = ragRes["data"].(string)
	}

	// 4. Save Assistant Reply to History
	s.repo.Create(&entities.ChatHistory{
		UserID:     userID,
		MealPlanID: req.MealPlanID,
		Role:       "assistant",
		Message:    response.Reply,
		Intent:     intent,
	})

	return &response, nil
}

func (s *ChatService) GetHistory(userID uuid.UUID) ([]dto.ChatHistoryResponse, error) {
	history, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var res []dto.ChatHistoryResponse
	for _, h := range history {
		res = append(res, dto.ChatHistoryResponse{
			ID:         h.ID,
			MealPlanID: h.MealPlanID,
			Role:       h.Role,
			Message:    h.Message,
			Intent:     h.Intent,
			CreatedAt:  h.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return res, nil
}
