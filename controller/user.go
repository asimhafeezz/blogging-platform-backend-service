package controller

import (
	"blogging-platform/backend-service/config"
	"blogging-platform/backend-service/model"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(ctx *gin.Context) {
	newUserData := ctx.Request.Body
	result, err := config.GetCollection("user").InsertOne(ctx, newUserData)
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
