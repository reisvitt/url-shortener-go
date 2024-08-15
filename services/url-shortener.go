package services

import "github.com/reisvitt/url-shortener-go/models"

func CreateShortenUrl(url *models.Url) (*models.Url, error) {
	// create a new shorten url
	// validate the shorten url
	// store the shorten url in the database
	// return the shortened url
	return nil, nil
}

func GetShortenedUrl(path *string) (*models.Url, error) {
	// get the full url from the database based on the shortened path
	// return the full url
	return nil, nil
}
