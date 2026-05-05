package user

import (
	"github.com/Mobilizes/materi-be-alpro/middlewares"
	authService "github.com/Mobilizes/materi-be-alpro/modules/auth/service"
	"github.com/Mobilizes/materi-be-alpro/modules/user/controller"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, ctrl *controller.UserController, jwtSvc *authService.JWTService) {
	users := r.Group("/user")
	{
		users.GET("/me", middlewares.Authentication(jwtSvc), ctrl.GetMe)
		users.PUT("/profile", middlewares.Authentication(jwtSvc), ctrl.UpdateProfile)
	}

	// Admin/Internal routes
	admin := r.Group("/admin/users")
	admin.Use(middlewares.Authentication(jwtSvc))
	{
		admin.GET("", ctrl.GetAllUsers)
		admin.GET("/:id", ctrl.GetUserByID)
	}
}
