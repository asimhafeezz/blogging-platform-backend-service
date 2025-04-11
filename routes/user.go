package routes

import (
	"blogging-platform/backend-service/controller"

	"github.com/gin-gonic/gin"
)

func userRoutes(r *gin.Engine) {
	userRoutes := r.Group("user")
	{
		userRoutes.POST("/", controller.RegisterUser)
		userRoutes.GET("/", controller.GetAllUsers)
		userRoutes.GET("/:id", controller.GetUserById)
	}
}
