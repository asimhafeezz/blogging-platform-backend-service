package controller

import (
	"blogging-platform/backend-service/config"
	"blogging-platform/backend-service/model"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(ctx *gin.Context) {
	var newUserData model.User
	if err := ctx.ShouldBindJSON(&newUserData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": true,
			"message": "Invalid JSON!",
		})
		return
	}
	result, err := config.GetCollection("user").InsertOne(ctx, newUserData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": false,
		"message": "user created successfully!",
		"data":    result,
	})
}

func GetAllUsers(ctx *gin.Context) {
	result, err := config.GetCollection("user").Find(ctx, bson.D{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": false,
		"message": "user created successfully!",
		"data":    result,
	})
}

func GetUserById(ctx *gin.Context) {
	userId := ctx.Params.ByName("id")
	result := config.GetCollection("iser").FindOne(ctx, bson.M{"_id": userId})
	if result.Err() != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": result.Err().Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": false,
		"message": "user created successfully!",
		"data":    result,
	})
}

func DeleteUserById(ctx *gin.Context) {
	userId := ctx.Params.ByName("id")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID can not be empty!",
		})
		return
	}

	_, err := config.GetCollection("user").DeleteOne(ctx, bson.M{"_id": userId})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": fmt.Sprintf("user with ID: %v, deleted successfully!", userId),
	})
}

// utils
func getUserById(ctx *gin.Context, id string) (*model.User, error) {
	var user model.User
	err := config.GetCollection("user").FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func getUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := config.GetCollection("user").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
