package main

import (
	"github.com/reisvitt/url-shortener-go/database"
	"github.com/reisvitt/url-shortener-go/repository"
	"github.com/reisvitt/url-shortener-go/router"
	"github.com/reisvitt/url-shortener-go/services"
)

func main() {
	// initialize database
	database.InitializeMongoDB("mongodb://admin:admin@localhost:27017/", "shortenerdb")

	// initialize respository
	repo := repository.NewMongoUrlRepository(database.Db)

	// initialize service
	service := services.NewUrlService(repo)

	// initialize router
	router.Initialize(service)
}
