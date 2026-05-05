package meal_plan

import (
	"github.com/Mobilizes/materi-be-alpro/middlewares"
	authService "github.com/Mobilizes/materi-be-alpro/modules/auth/service"
	"github.com/Mobilizes/materi-be-alpro/modules/meal_plan/controller"
	"github.com/gin-gonic/gin"
)

func RegisterMealPlanRoutes(r *gin.RouterGroup, ctrl *controller.MealPlanController, jwtSvc *authService.JWTService) {
	plans := r.Group("/meal-plan")
	plans.Use(middlewares.Authentication(jwtSvc))
	{
		plans.POST("/generate", ctrl.Generate)
		plans.GET("", ctrl.GetAll)
		plans.GET("/:id", ctrl.GetByID)
		plans.GET("/:id/versions", ctrl.GetVersions)
		plans.DELETE("/:id", ctrl.Delete)
	}
}
