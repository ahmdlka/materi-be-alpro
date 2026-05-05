package service

import "github.com/google/uuid"

const (
	IntentRefineMenu  = "REFINE_MENU"
	IntentAskQuestion = "ASK_QUESTION"
)

type IntentService struct{}

func NewIntentService() *IntentService {
	return &IntentService{}
}

func (s *IntentService) Classify(mealPlanID *uuid.UUID) string {
	if mealPlanID != nil && *mealPlanID != uuid.Nil {
		return IntentRefineMenu
	}
	return IntentAskQuestion
}
