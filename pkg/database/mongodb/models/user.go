// pkg/controllers/models/user.go
package models

// User represents the model for a user.
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
