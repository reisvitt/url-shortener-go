package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reisvitt/url-shortener-go/dto"
	"github.com/reisvitt/url-shortener-go/utils"
)

func (h *UrlHandler) CreateShortenUrlHandler(ctx *gin.Context) {
	urlDTO := dto.UrlDTO{}

	if err := ctx.BindJSON(&urlDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"status":  "error",
			"error":   err.Error(),
		})
		return
	}

	if err := urlDTO.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"status":  "error",
			"error":   err.Error(),
		})
		return
	}

	// create a new shorten url
	url, err := h.service.CreateShortenUrl(&urlDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create shorten URL",
			"status":  "error",
			"error":   err.Error(),
		})
		return
	}

	domain := *utils.GetDomain(ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"url": fmt.Sprintf("%s/%s", domain, url.ID),
	})
}
