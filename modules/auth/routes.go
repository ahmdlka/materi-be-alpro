package auth

import (
	"github.com/Mobilizes/materi-be-alpro/modules/auth/controller"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.RouterGroup, ctrl *controller.AuthController) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", ctrl.Register)
		auth.POST("/login", ctrl.Login)
	}
}
