package user

import (
	"github.com/Mobilizes/materi-be-alpro/middlewares"
	authService "github.com/Mobilizes/materi-be-alpro/modules/auth/service"
	"github.com/Mobilizes/materi-be-alpro/modules/user/controller"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, ctrl *controller.UserController, jwtSvc *authService.JWTService) {
	users := r.Group("/users")
	{
		users.POST("", ctrl.CreateUser)                                         // POST /api/users
		users.GET("", middlewares.Authentication(jwtSvc), ctrl.GetAllUsers)     // GET  /api/users (protected)
		users.GET("/:id", middlewares.Authentication(jwtSvc), ctrl.GetUserByID) // GET  /api/users/:id (protected)
	}
}
