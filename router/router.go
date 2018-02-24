package router

import (
	"github.com/gin-gonic/gin"
	"auth2/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("auth-api/v1")
	// operation
	v1.GET("operation", controllers.GetOperations)
	v1.GET("operation/:id", controllers.GetOperation)
	v1.POST("operation", controllers.CreateOperation)
	v1.PUT("operation/:id", controllers.UpdateOperation)

	return router
}
