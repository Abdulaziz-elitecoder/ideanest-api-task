// pkg/api/middleware/auth_middleware.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware function to check for valid authentication.
func AuthMiddleware(c *gin.Context) {
	// Retrieve the token from the request header, query parameter, or cookie
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// TODO: Perform token validation and verification logic here
	// You might want to check the token against a stored list of valid tokens

	// Example: Check if the token is valid (replace this with your actual logic)
	if token != "valid_token" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Continue processing the request
	c.Next()
}
