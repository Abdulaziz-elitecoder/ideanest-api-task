// utils/app.go
package utils

import (
	"github.com/Abdulaziz-elitecoder/ideanest-api-task/pkg/api/routes"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures and returns the Gin router for your application.
func SetupRouter() *gin.Engine {
	// Create a new Gin router
	router := gin.Default()

	// Set up your API routes
	routes.SetupRoutes(router)

	return router
}
