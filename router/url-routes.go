package router

import (
	"github.com/gin-gonic/gin"
	"github.com/reisvitt/url-shortener-go/handlers"
	"github.com/reisvitt/url-shortener-go/services"
)

func initializeUrlRoutes(router *gin.Engine, service *services.UrlService) {

	// defining API group with version prefix
	v1 := router.Group("")

	// initializing URL handlers
	urlHandlers := handlers.NewUrlHandler(service)

	// defining route
	v1.POST("/shorten-url", urlHandlers.CreateShortenUrlHandler)
	v1.GET("/:path", urlHandlers.GetShortedUrlHandler)
}
