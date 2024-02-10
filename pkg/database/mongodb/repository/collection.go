// pkg/controllers/database/mongodb/repository/collection.go
package repository

import (
	"context"

	"github.com/Abdulaziz-elitecoder/ideanest-api-task/pkg/database/mongodb"
	"github.com/Abdulaziz-elitecoder/ideanest-api-task/pkg/database/mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection

func init() {
	UserCollection = mongodb.Client.Database("your-database-name").Collection("users")
}

// InsertUser inserts a new user into MongoDB.
func InsertUser(user models.User) error {
	_, err := UserCollection.InsertOne(context.Background(), user)
	return err
}

// GetUserByEmail retrieves a user from MongoDB based on email.
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := UserCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	return &user, err
}
