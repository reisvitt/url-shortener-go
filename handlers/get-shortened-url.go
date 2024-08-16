package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UrlHandler) GetShortedUrlHandler(ctx *gin.Context) {
	uri := ctx.Request.RequestURI
	uri = uri[1:] // remove "/" from start of url

	fullUrl, err := h.service.GetShortenedUrl(&uri)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "URL not found",
			"status":  "not-found",
		})
		return
	}

	ctx.Redirect(http.StatusPermanentRedirect, *fullUrl)
}
