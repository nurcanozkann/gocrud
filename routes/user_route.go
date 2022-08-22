package routes

import (
	"cihanozhan.com/dbgo/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	userController := controllers.NewUserController()
	router.GET("/users", userController.GetAllUsers)
	router.GET("/users/:id", userController.GetUserById)
	router.POST("/users", userController.CreateUser)
	router.DELETE("/users/:id", userController.DeleteUser)
	router.PUT("/users/:id", userController.UpdateUser)
}
