package main

import (
	"blogging-platform/backend-service/config"
	"blogging-platform/backend-service/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// load envs
	config.LoadEnv()

	// connect with mongodb
	config.ConnectDB()

	// router
	router := gin.Default()

	router.Use(cors.Default())

	// register all the routes
	routes.RegisterRoutes(router)

	port := config.GetEnv("PORT", "8080")
	router.Run(":" + port)

}
