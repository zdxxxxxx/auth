package router

import (
	"github.com/gin-gonic/gin"
	"auth2/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("auth-api/v1")
	v1.GET("operations", controllers.GetOperations)
	v1.GET("operations/:id", controllers.GetOperation)

	return router
}
