package validation

import (
    "github.com/gin-gonic/gin"
    "github.com/Mobilizes/materi-be-alpro/modules/user/dto"
)

func ValidateCreateUser(c *gin.Context) (*dto.CreateUserRequest, error) {
    var req dto.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        return nil, err
    }
    return &req, nil
}
