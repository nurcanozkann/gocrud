package controllers

import (
	"net/http"
	"strconv"

	"cihanozhan.com/dbgo/common"
	"cihanozhan.com/dbgo/models"
	"cihanozhan.com/dbgo/service"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	GetAllUsers(c *gin.Context)
	GetUserById(c *gin.Context)
	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

func NewUserController() *userController {
	return &userController{
		userService: service.NewUserService(),
	}
}

func (controller *userController) GetAllUsers(c *gin.Context) {
	getUsers, restErr := controller.userService.GetAllUsers()
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, getUsers)
}

func (controller *userController) GetUserById(c *gin.Context) {
	id, errParse := strconv.Atoi(c.Param("id"))
	if errParse != nil {
		restErr := common.BadRequest("Request body or parameters wrong.")
		c.JSON(restErr.Status, restErr)
		return
	}
	getUser, restErr := controller.userService.GetUserById(id)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, getUser)
}

func (controller *userController) DeleteUser(c *gin.Context) {
	id, errParse := strconv.Atoi(c.Param("id"))
	if errParse != nil {
		restErr := common.BadRequest("Request body or parameters wrong.")
		c.JSON(restErr.Status, restErr)
		return
	}

	restErr := controller.userService.DeleteUser(id)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"isRemoved": true})
}

func (controller *userController) CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		restErr := common.BadRequest("Request body or parameters wrong.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := controller.userService.CreateUser(&newUser)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (controller *userController) UpdateUser(c *gin.Context) {
	var user models.User
	id, errParse := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&user)

	if errParse != nil && err != nil {
		restErr := common.BadRequest("Request body or parameters wrong.")
		c.JSON(restErr.Status, restErr)
		return
	}

	updateUser, restErr := controller.userService.UpdateUser(id, &user)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, updateUser)
}
