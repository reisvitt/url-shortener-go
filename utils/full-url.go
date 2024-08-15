package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetDomain(ctx *gin.Context) *string {
	// Obtendo o protocolo (normalmente pode ser inferido se você estiver atrás de um proxy)
	protocol := "http"
	if ctx.Request.TLS != nil {
		protocol = "https"
	}

	// Montando a URL completa
	domain := fmt.Sprintf("%s://%s", protocol, ctx.Request.Host)

	return &domain
}
