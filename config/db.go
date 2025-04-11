package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var mongoDBClient *mongo.Client
var db *mongo.Database

func ConnectDB() {
	mongoURI := GetEnv("MONGO_URI", "mongodb://bloggin-platform-mongo:27017/blogging-platform-local")
	println("MONGO URI :: ", mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	mongoDBClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		println("MongoDB is not connected!")
	}
	println("MongoDB is connected!")
	db = mongoDBClient.Database("blogging-platform-local")
}

func GetCollection(collection string) *mongo.Collection {
	return db.Collection(collection)
}
