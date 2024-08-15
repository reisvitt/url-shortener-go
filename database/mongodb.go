package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func InitializeMongoDB(uri string) {
	// initialize MongoDB client with provided URI
	// establish connection to the MongoDB server
	// return error if connection fails

	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Verifica a conexão
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	MongoClient = client

	print("Connected to MongoDB!")
}
