package controllers

import (
	"net/http"

	"strconv"

	"cihanozhan.com/dbgo/common"
	"cihanozhan.com/dbgo/config"
	"cihanozhan.com/dbgo/models"
	"cihanozhan.com/dbgo/service"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func GetAllUsers(c *gin.Context) {
	getUsers, restErr := service.GetAllUsers()
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, getUsers)
}

func GetUserByIds(c *gin.Context) {
	var user models.User
	config.DB.Find(&user, c.Param("id"))
	c.JSON(200, &user)
}

func GetUserById(c *gin.Context) {
	id, errParse := strconv.Atoi(c.Param("id"))
	if errParse != nil {
		restErr := common.BadRequest("Request body or parameters wrong.")
		c.JSON(restErr.Status, restErr)
		return
	}
	getUser, restErr := service.GetUserById(id)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, getUser)
}

func DeleteUser(c *gin.Context) {
	id, errParse := strconv.Atoi(c.Param("id"))
	if errParse != nil {
		restErr := common.BadRequest("Request body or parameters wrong.")
		c.JSON(restErr.Status, restErr)
		return
	}

	restErr := service.DeleteUser(id)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"isRemoved": true})
}

func CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		restErr := common.BadRequest("Request body or parameters wrong.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := service.CreateUser(&newUser)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id, errParse := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&user)

	if errParse != nil && err != nil {
		restErr := common.BadRequest("Request body or parameters wrong.")
		c.JSON(restErr.Status, restErr)
		return
	}

	updateUser, restErr := service.UpdateUser(id, &user)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, updateUser)
}
