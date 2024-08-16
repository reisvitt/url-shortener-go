package services

import (
	"github.com/reisvitt/url-shortener-go/dto"
	"github.com/reisvitt/url-shortener-go/models"
	"github.com/reisvitt/url-shortener-go/repository"
	"github.com/reisvitt/url-shortener-go/utils"
)

type UrlService struct {
	repo repository.UrlRepository
}

func NewUrlService(repository repository.UrlRepository) *UrlService {
	return &UrlService{repo: repository}
}

func (s *UrlService) CreateShortenUrl(urlDTO *dto.UrlDTO) (*models.Url, error) {

	url := &models.Url{
		FullUrl: urlDTO.Url,
	}

	var shortenUrl string

	for {
		// create a new shorten url
		random, err := utils.GerenateRandomPath()
		if err != nil {
			return nil, err
		}

		shortenUrl = random

		// validate the shorten url
		has, _ := s.repo.GetByID(&shortenUrl)
		if has == nil {
			break
		}
	}

	// store the shorten url in the database
	url.ID = shortenUrl
	dabatabaseUrl, err := s.repo.Create(url)

	if err != nil {
		return nil, err
	}

	// return the shortened url
	return dabatabaseUrl, nil
}

func (s *UrlService) GetShortenedUrl(path *string) (*string, error) {
	// get the full url from the database based on the shortened path
	url, err := s.repo.GetByID(path)

	if err != nil {
		return nil, err
	}

	// return the full url
	return &url.FullUrl, nil
}
