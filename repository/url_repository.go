package repository

import (
	"context"
	"time"

	"github.com/reisvitt/url-shortener-go/configs"
	"github.com/reisvitt/url-shortener-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUrlRepository struct {
	collection *mongo.Collection
}

func NewMongoUrlRepository(db *mongo.Database) UrlRepository {
	return &mongoUrlRepository{collection: db.Collection("urls")}
}

func (repo *mongoUrlRepository) Create(url *models.Url) (*models.Url, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if url.CreatedAt.IsZero() {
		url.CreatedAt = time.Now()
	}

	if url.ExpiresAt.IsZero() {
		url.ExpiresAt = url.CreatedAt.Add(configs.TIME_TO_EXPIRES)
	}

	_, err := repo.collection.InsertOne(ctx, url)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (repo *mongoUrlRepository) GetByID(id *string) (*models.Url, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var url models.Url
	err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&url)
	if err != nil {
		return nil, err
	}

	return &url, nil
}
