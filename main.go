package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/reisvitt/url-shortener-go/database"
	"github.com/reisvitt/url-shortener-go/repository"
	"github.com/reisvitt/url-shortener-go/router"
	"github.com/reisvitt/url-shortener-go/services"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	dbConnection := os.Getenv("DB_CONNECTION")
	dbName := os.Getenv("DB_NAME")

	// initialize database
	database.InitializeMongoDB(dbConnection, dbName)

	// create TTL index on 'created-at' field in 'urls' collection
	database.CreateTTLIndex(database.Db.Collection("urls"))

	// initialize respository
	repo := repository.NewMongoUrlRepository(database.Db)

	// initialize service
	service := services.NewUrlService(repo)

	// initialize router
	router.Initialize(service)
}
