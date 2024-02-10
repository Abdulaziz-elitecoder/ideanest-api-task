// pkg/api/handlers/auth_handlers.go
package handlers

import (
	"net/http"

	"github.com/Abdulaziz-elitecoder/ideanest-api-task/pkg/controllers"
	"github.com/Abdulaziz-elitecoder/ideanest-api-task/pkg/database/mongodb/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Signup handles user registration.
func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the user's password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Call the controller to handle user registration
	err = controllers.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Signin handles user authentication.
func Signin(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the controller to handle user authentication
	user, err := controllers.AuthenticateUser(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// TODO: Generate and return access and refresh tokens
	// Implement your token generation logic here

	c.JSON(http.StatusOK, gin.H{"message": "Signin successful"})
}

// RefreshToken handles token refreshing.
func RefreshToken(c *gin.Context) {
	var refreshTokenRequest struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&refreshTokenRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement token refreshing logic
	// Call the controller to handle token refreshing

	c.JSON(http.StatusOK, gin.H{"message": "Token refreshed successfully"})
}
