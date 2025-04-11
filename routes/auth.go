package routes

import (
	"blogging-platform/backend-service/controller"

	"github.com/gin-gonic/gin"
)

func authRoutes(router *gin.Engine) {
	authRoutes := router.Group("auth")
	{
		authRoutes.POST("/register", controller.RegisterUser)
		authRoutes.POST("/login", controller.LoginUser)
	}
}
