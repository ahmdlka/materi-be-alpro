package chat

import (
	"github.com/Mobilizes/materi-be-alpro/middlewares"
	authService "github.com/Mobilizes/materi-be-alpro/modules/auth/service"
	"github.com/Mobilizes/materi-be-alpro/modules/chat/controller"
	"github.com/gin-gonic/gin"
)

func RegisterChatRoutes(r *gin.RouterGroup, ctrl *controller.ChatController, jwtSvc *authService.JWTService) {
	chats := r.Group("/chat")
	chats.Use(middlewares.Authentication(jwtSvc))
	{
		chats.POST("", ctrl.SendMessage)
		chats.GET("/history", ctrl.GetHistory)
	}
}
