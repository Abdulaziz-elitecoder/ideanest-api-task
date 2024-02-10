// pkg/controllers/controllers.go
package controllers

import (
	"github.com/Abdulaziz-elitecoder/ideanest-api-task/pkg/database/mongodb/models"
	"github.com/Abdulaziz-elitecoder/ideanest-api-task/pkg/database/mongodb/repository"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser registers a new user.
func RegisterUser(user models.User) error {
	// Call the repository to insert the user into MongoDB
	return repository.InsertUser(user)
}

// AuthenticateUser authenticates a user based on credentials.
func AuthenticateUser(email, password string) (*models.User, error) {
	// Call the repository to retrieve the user from MongoDB based on email
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// Compare the provided password with the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
