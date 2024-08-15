package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateTTLIndex(colletionName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"created-at": 1},                  // Field that will be used for the TTL index
		Options: options.Index().SetExpireAfterSeconds(0), // Expires exactly on the date set
	}

	collection := GetDbCollection(colletionName) // collection that will be used for the TTL index

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		fmt.Println("Failed to create TTL index:", err)
		return err
	}

	fmt.Println("TTL index created successfully")
	return nil
}
