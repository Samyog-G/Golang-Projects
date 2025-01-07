package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env") //attempts to load environment variables from .env file into the environment
	if err != nil {
		log.Fatal("Error loading .env file") //if therre is error in loading the .env file
	}
	MongoDB := os.Getenv("MONGODB_URL") //retrieves MongoDB connection URL

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB)) //creates a new mongodb client
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //creates new context with timeout of 10 sec
	defer cancel()
	err = client.Connect(ctx) //connection to mongodb server using the client and context
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	return client

}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName) //retrieves specified collection from database named "cluster0"
	return collection
}
