package router

import (
	"github.com/gin-gonic/gin"
	"auth2/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("auth-api/v1")
	// 操作
	v1.GET("operation", controllers.GetOperations)
	v1.GET("operation/:id", controllers.GetOperation)
	v1.POST("operation", controllers.CreateOperation)
	v1.PUT("operation/:id", controllers.UpdateOperation)
	v1.DELETE("operation/:id", controllers.DeleteOperation)

	// 平台
	v1.GET("app", controllers.GetApps)
	v1.GET("app/:id", controllers.GetApp)
	v1.POST("app", controllers.CreateApp)
	v1.PUT("app/:id", controllers.UpdateApp)
	v1.DELETE("app/:id", controllers.DeleteApp)
	v1.PATCH("app/:id", controllers.UpdateAppContent)

	// 权限
	v1.POST("auth", controllers.CreateAuth)
	v1.PUT("auth/:id", controllers.UpdateAuth)
	v1.DELETE("auth/:id", controllers.DeleteAuth)
	v1.GET("auth/:id", controllers.GetAuth)
	v1.GET("auths/:id", controllers.GetAuths)
	// 授权
	v1.POST("authorize", controllers.CreateUserAuth)
	v1.GET("authorize", controllers.GetUserAuth)
	v1.DELETE("authorize", controllers.DeleteUserAuth)
	// 鉴权
	v1.POST("check", controllers.CheckAuth)
	return router
}
