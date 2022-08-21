package routes

import (
	"cihanozhan.com/dbgo/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:id", controllers.GetUserById)
	router.POST("/users", controllers.CreateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
	router.PUT("/users/:id", controllers.UpdateUser)
}
