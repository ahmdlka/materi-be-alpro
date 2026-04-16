package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/Mobilizes/materi-be-alpro/modules/auth/controller"
)

func RegisterAuthRoutes(r *gin.RouterGroup, ctrl *controller.AuthController) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", ctrl.Login)
	}
}
