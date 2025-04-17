package routes

import (
	"blogging-platform/backend-service/controller"
	"blogging-platform/backend-service/middleware"

	"github.com/gin-gonic/gin"
)

func userRoutes(r *gin.Engine) {
	userRoutes := r.Group("user")
	userRoutes.Use(middleware.VerifyAccessToken())
	{
		userRoutes.POST("/", controller.RegisterUser)
		userRoutes.GET("/", controller.GetAllUsers)
		userRoutes.GET("/:id", controller.GetUserById)
		userRoutes.DELETE("/:id", controller.DeleteUserById)
	}
}
