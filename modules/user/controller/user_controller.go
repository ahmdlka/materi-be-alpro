package controller

import (
	"net/http"
	"strconv"

	"github.com/Mobilizes/materi-be-alpro/modules/user/service"
	"github.com/Mobilizes/materi-be-alpro/modules/user/validation"
	"github.com/Mobilizes/materi-be-alpro/pkg/utils"
	"github.com/gin-gonic/gin"
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

func (ctrl *UserController) GetUser(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Param("id"))

	profile, err := ctrl.service.Get(
		c.Request.Context(),
		userID,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"success": false,
				"message": "failed to get user profile",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": "User profile retrieved successfully",
			"data": gin.H{
				"id":       profile.ID,
				"name":     profile.Name,
				"email":    profile.Email,
				"password": profile.Password,
			},
		},
	)
}

func (ctrl *UserController) GetAllUser(c *gin.Context) {

	// Mengambil properti berdasarkan filter
	data, err := ctrl.service.GetAll(
		c.Request.Context(),
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)
		return
	}

	// Response sukses dengan metadata pagination
	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message": "get property data with filter and pagination",
			"data":    data,
		},
	)
}
