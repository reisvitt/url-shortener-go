package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	// defining API group with version prefix
	v1 := router.Group("/api/v1")

	// defining route
	v1.POST("/shorten-url", handlerCreateShortener)
	v1.GET("/:path", handlerGetShortedUrl)
}

func handlerCreateShortener(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "URL shortening request received",
		"status":  "success",
	})
}

func handlerGetShortedUrl(ctx *gin.Context) {
	ctx.JSON(http.StatusPermanentRedirect, gin.H{
		"message": "Redirecting to path",
		"status":  "success",
	})
}
