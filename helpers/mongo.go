package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	UsersCollection          *mongo.Collection
	ProjectsCollection       *mongo.Collection
	TasksCollection          *mongo.Collection
	CommentsCollection       *mongo.Collection
	CommentsBucketCollection *mongo.Collection
	Ctx                      = context.TODO()
)

func MongoConnection() {
	// Set client options
	// host := "localhost"
	// port := "27017"
	// clientOptions := options.Client().ApplyURI("mongodb://" + host + ":" + port)
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.mvebh.mongodb.net")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	db := client.Database("todo")

	ProjectsCollection = db.Collection("projects")
	UsersCollection = db.Collection("users")
	TasksCollection = db.Collection("tasks")
	CommentsCollection = db.Collection("comments")
	CommentsBucketCollection = db.Collection("commentsBucket")

}
