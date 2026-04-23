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

func (ctrl *UserController) GetUserByID(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Param("id"))

	user, err := ctrl.service.Get(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Gagal mendapatkan profil user")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Profil user berhasil didapatkan", user)
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {

	users, err := ctrl.service.GetAll()
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Gagal mendapatkan profil users")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Profil user berhasil didapatkan", users)
}
