package validation

import (
	"github.com/Mobilizes/materi-be-alpro/modules/chat/dto"
	"github.com/gin-gonic/gin"
)

func ValidateChat(c *gin.Context) (*dto.ChatRequest, error) {
	var req dto.ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
