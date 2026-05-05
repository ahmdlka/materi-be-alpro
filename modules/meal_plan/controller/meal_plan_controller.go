package controller

import (
	"net/http"

	"github.com/Mobilizes/materi-be-alpro/modules/meal_plan/service"
	"github.com/Mobilizes/materi-be-alpro/modules/meal_plan/validation"
	"github.com/Mobilizes/materi-be-alpro/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MealPlanController struct {
	service *service.MealPlanService
}

func NewMealPlanController(service *service.MealPlanService) *MealPlanController {
	return &MealPlanController{service: service}
}

func (ctrl *MealPlanController) Generate(c *gin.Context) {
	userID, _ := c.Get("user_id")

	req, err := validation.ValidateGenerateMealPlan(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	plan, err := ctrl.service.Generate(userID.(uuid.UUID), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Meal plan generated successfully", plan)
}

func (ctrl *MealPlanController) GetAll(c *gin.Context) {
	userID, _ := c.Get("user_id")

	plans, err := ctrl.service.GetByUserID(userID.(uuid.UUID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch meal plans")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Meal plans fetched successfully", plans)
}

func (ctrl *MealPlanController) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	plan, err := ctrl.service.GetByID(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Meal plan not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Meal plan fetched successfully", plan)
}

func (ctrl *MealPlanController) GetVersions(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	versions, err := ctrl.service.GetVersions(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch versions")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Versions fetched successfully", versions)
}

func (ctrl *MealPlanController) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	if err := ctrl.service.Delete(id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete meal plan")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Meal plan deleted successfully", nil)
}
