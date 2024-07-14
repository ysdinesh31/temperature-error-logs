package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectMongoDB() error {
	var err error
	var mongo_uri = os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(mongo_uri)
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return err
	}

	// Check the connection
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
		return err
	}

	return nil
}
