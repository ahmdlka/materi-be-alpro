package validation

import (
	"github.com/Mobilizes/materi-be-alpro/modules/meal_plan/dto"
	"github.com/gin-gonic/gin"
)

func ValidateGenerateMealPlan(c *gin.Context) (*dto.GenerateMealPlanRequest, error) {
	var req dto.GenerateMealPlanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
