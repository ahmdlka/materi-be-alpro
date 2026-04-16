package validation

import (
	"github.com/gin-gonic/gin"
	"github.com/Mobilizes/materi-be-alpro/modules/auth/dto"
)

func ValidateLogin(c *gin.Context) (*dto.LoginRequest, error) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
