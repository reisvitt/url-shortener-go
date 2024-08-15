package main

import (
	"github.com/reisvitt/url-shortener-go/database"
	"github.com/reisvitt/url-shortener-go/router"
)

func main() {
	// initialize database
	database.InitializeMongoDB("mongodb://admin:admin@localhost:27017/")

	// initialize router
	router.Initialize()
}
