package router

import (
	"github.com/gin-gonic/gin"
	"github.com/reisvitt/url-shortener-go/handlers"
)

func initializeRoutes(router *gin.Engine) {

	// defining API group with version prefix
	v1 := router.Group("/api/v1")

	// defining route
	v1.POST("/shorten-url", handlers.CreateShortenUrlHandler)
	v1.GET("/:path", handlers.GetShortedUrlHandler)
}
