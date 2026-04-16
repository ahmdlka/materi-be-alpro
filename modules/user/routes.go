package user

import (
	authService "github.com/Mobilizes/materi-be-alpro/modules/auth/service"
	"github.com/Mobilizes/materi-be-alpro/modules/user/controller"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, ctrl *controller.UserController, jwtSvc *authService.JWTService) {
	users := r.Group("/users")
	{
		users.POST("", ctrl.CreateUser)
		users.GET("/", ctrl.GetUser)
		users.GET("/:id", ctrl.GetAllUser)
	}
}
