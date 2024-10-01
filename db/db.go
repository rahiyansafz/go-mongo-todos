package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var collection *mongo.Collection

func Connect() error {
	uri := os.Getenv("MONGO_URI")
	username := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	password := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")

	if uri == "" || username == "" || password == "" {
		return fmt.Errorf("missing required environment variables")
	}

	log.Printf("Connecting to MongoDB with URI: %s", uri)

	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the database
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB successfully")

	collection = client.Database("todos").Collection("todos")
	log.Println("Collection 'todos' in database 'todos' is ready")

	return nil
}

func GetCollection() *mongo.Collection {
	return collection
}
