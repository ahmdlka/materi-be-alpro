package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Mobilizes/materi-be-alpro/modules/auth/dto"
	"github.com/Mobilizes/materi-be-alpro/modules/auth/service"
	"github.com/Mobilizes/materi-be-alpro/modules/auth/validation"
	"github.com/Mobilizes/materi-be-alpro/pkg/utils"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ctrl *AuthController) Login(c *gin.Context) {
	req, err := validation.ValidateLogin(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := ctrl.authService.Login(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Login berhasil", dto.TokenResponse{Token: token})
}
