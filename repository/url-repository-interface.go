package repository

import "github.com/reisvitt/url-shortener-go/models"

type UrlRepository interface {
	Create(url *models.Url) (*models.Url, error)
	GetByID(id *string) (*models.Url, error)
}
