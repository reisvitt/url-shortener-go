package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

func InitializeMongoDB(uri, database string) {
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

	// verify connection to MongoDB server
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	Db = client.Database(database)

	fmt.Println("Connected to MongoDB!")
}
