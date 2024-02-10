// pkg/api/routes/routes.go
package routes

import (
	"github.com/Abdulaziz-elitecoder/ideanest-api-task/pkg/api/handlers"
	"github.com/Abdulaziz-elitecoder/ideanest-api-task/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes and sets up all the routes for the application.
func SetupRoutes(router *gin.Engine) {
	// Set up routes for authentication handlers
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", handlers.Signup)
		authGroup.POST("/signin", handlers.Signin)
		authGroup.POST("/refresh-token", handlers.RefreshToken)
	}

	// Set up routes for organization handlers (example)
	orgGroup := router.Group("/organization")
	orgGroup.Use(middleware.AuthMiddleware) // Add authentication middleware as needed
	{
		orgGroup.POST("/", handlers.CreateOrganization)
		orgGroup.GET("/:organization_id", handlers.ReadOrganization)
		orgGroup.PUT("/:organization_id", handlers.UpdateOrganization)
		orgGroup.DELETE("/:organization_id", handlers.DeleteOrganization)
		orgGroup.POST("/:organization_id/invite", handlers.InviteUserToOrganization)
	}
}
