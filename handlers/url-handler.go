package handlers

import "github.com/reisvitt/url-shortener-go/services"

type UrlHandler struct {
	service *services.UrlService
}

func NewUrlHandler(service *services.UrlService) *UrlHandler {
	return &UrlHandler{
		service: service,
	}
}
