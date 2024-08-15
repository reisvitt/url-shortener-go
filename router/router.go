package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// initialize gin router with default configuration
	r := gin.Default()

	// defining route
	initializeRoutes(r)

	// listening and serving on :8080
	r.Run()
}
