package app

import (
	"cihanozhan.com/dbgo/config"
	"cihanozhan.com/dbgo/routes"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.New()
	config.ConnDB()
	routes.UserRoute(router)
	router.Run(":8080")
}
