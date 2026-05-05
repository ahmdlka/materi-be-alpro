package controller

import (
	"net/http"

	"github.com/Mobilizes/materi-be-alpro/modules/chat/service"
	"github.com/Mobilizes/materi-be-alpro/modules/chat/validation"
	"github.com/Mobilizes/materi-be-alpro/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ChatController struct {
	service *service.ChatService
}

func NewChatController(service *service.ChatService) *ChatController {
	return &ChatController{service: service}
}

func (ctrl *ChatController) SendMessage(c *gin.Context) {
	userID, _ := c.Get("user_id")

	req, err := validation.ValidateChat(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := ctrl.service.ProcessMessage(userID.(uuid.UUID), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Pesan berhasil diproses", response)
}

func (ctrl *ChatController) GetHistory(c *gin.Context) {
	userID, _ := c.Get("user_id")

	history, err := ctrl.service.GetHistory(userID.(uuid.UUID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengambil riwayat chat")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Riwayat chat berhasil diambil", history)
}
