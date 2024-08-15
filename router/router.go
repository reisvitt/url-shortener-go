package router

import (
	"github.com/gin-gonic/gin"
	"github.com/reisvitt/url-shortener-go/services"
)

func Initialize(service *services.UrlService) {
	// initialize gin router with default configuration
	r := gin.Default()

	// defining route
	initializeUrlRoutes(r, service)

	// listening and serving on :8080
	r.Run()
}
