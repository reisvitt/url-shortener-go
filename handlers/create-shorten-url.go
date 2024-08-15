package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UrlHandler) CreateShortenUrlHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "URL shortening request received",
		"status":  "success",
	})
}
