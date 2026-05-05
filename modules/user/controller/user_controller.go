package controller

import (
	"net/http"

	"github.com/Mobilizes/materi-be-alpro/modules/user/service"
	"github.com/Mobilizes/materi-be-alpro/modules/user/validation"
	"github.com/Mobilizes/materi-be-alpro/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	req, err := validation.ValidateCreateUser(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := ctrl.service.CreateUser(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal membuat user")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "User berhasil dibuat", user)
}

func (ctrl *UserController) GetUserByID(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid User ID format")
		return
	}

	user, err := ctrl.service.Get(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Gagal mendapatkan profil user")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Profil user berhasil didapatkan", user)
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.service.GetAll()
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Gagal mendapatkan profil users")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Profil user berhasil didapatkan", users)
}

func (ctrl *UserController) GetMe(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, err := ctrl.service.Get(userID.(uuid.UUID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User tidak ditemukan")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Data user berhasil diambil", user)
}

func (ctrl *UserController) UpdateProfile(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	req, err := validation.ValidateUpdateProfile(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := ctrl.service.UpdateProfile(userID.(uuid.UUID), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal update profil")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Profil berhasil diupdate", user)
}
