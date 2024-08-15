package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetShortedUrlHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusPermanentRedirect, gin.H{
		"message": "Redirecting to path",
		"status":  "success",
	})
}
