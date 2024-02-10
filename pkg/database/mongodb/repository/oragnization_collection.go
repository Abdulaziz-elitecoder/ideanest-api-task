// pkg/controllers/database/mongodb/repository/organization_collection.go
package repository

import (
	"context"
	"errors"

	"github.com/Abdulaziz-elitecoder/ideanest-api-task/pkg/database/mongodb"
	"github.com/Abdulaziz-elitecoder/ideanest-api-task/pkg/database/mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var OrganizationCollection *mongo.Collection

func init() {
	OrganizationCollection = mongodb.Client.Database("your-database-name").Collection("organizations")
}

// InsertOrganization inserts a new organization into MongoDB.
func InsertOrganization(org models.Organization) (string, error) {
	result, err := OrganizationCollection.InsertOne(context.Background(), org)
	if err != nil {
		return "", err
	}

	// Retrieve the inserted organization's ID
	orgID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("Failed to get the inserted organization ID")
	}

	return orgID.Hex(), nil
}

// GetOrganization retrieves an organization from MongoDB based on its ID.
func GetOrganization(orgID string) (*models.Organization, error) {
	objectID, err := primitive.ObjectIDFromHex(orgID)
	if err != nil {
		return nil, errors.New("Invalid organization ID")
	}

	var org models.Organization
	err = OrganizationCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&org)
	if err != nil {
		return nil, err
	}

	return &org, nil
}

// UpdateOrganization updates an organization in MongoDB.
func UpdateOrganization(orgID string, updatedOrg models.Organization) error {
	objectID, err := primitive.ObjectIDFromHex(orgID)
	if err != nil {
		return errors.New("Invalid organization ID")
	}

	_, err = OrganizationCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": objectID},
		bson.M{"$set": updatedOrg},
	)
	if err != nil {
		return err
	}

	return nil
}

// DeleteOrganization deletes an organization from MongoDB.
func DeleteOrganization(orgID string) error {
	objectID, err := primitive.ObjectIDFromHex(orgID)
	if err != nil {
		return errors.New("Invalid organization ID")
	}

	_, err = OrganizationCollection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	return nil
}
