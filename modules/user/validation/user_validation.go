package validation

import (
	"github.com/Mobilizes/materi-be-alpro/modules/user/dto"
	"github.com/gin-gonic/gin"
)

func ValidateCreateUser(c *gin.Context) (*dto.CreateUserRequest, error) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func ValidateUpdateProfile(c *gin.Context) (*dto.UpdateProfileRequest, error) {
	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
