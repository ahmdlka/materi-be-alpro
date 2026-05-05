package validation

import (
	"github.com/Mobilizes/materi-be-alpro/modules/auth/dto"
	userDto "github.com/Mobilizes/materi-be-alpro/modules/user/dto"
	"github.com/gin-gonic/gin"
)

func ValidateLogin(c *gin.Context) (*dto.LoginRequest, error) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func ValidateRegister(c *gin.Context) (*userDto.CreateUserRequest, error) {
	var req userDto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
