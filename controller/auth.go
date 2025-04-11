package controller

import (
	"blogging-platform/backend-service/config"
	"blogging-platform/backend-service/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUser(c *gin.Context) {
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid JSON",
		})
		return
	}

	userByEmail, _ := getUserByEmail(c.Request.Context(), data.Email)

	if userByEmail != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "user already exists",
		})
		return
	}

	updatedUser, err := config.GetCollection("user").InsertOne(c, data)
	data.Id = updatedUser.InsertedID.(primitive.ObjectID)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "user registered successfully!",
		"data":    data,
	})
}

func LoginUser(ctx *gin.Context) {
	var data model.User

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "email or password is incorrect!",
		})
		return
	}

	userByEmail, err := getUserByEmail(ctx.Request.Context(), data.Email)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "email or password is not correct!",
		})
		return
	}

	if userByEmail.Password != data.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "email or password is not correct!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "user logged in successfully!",
	})
}
