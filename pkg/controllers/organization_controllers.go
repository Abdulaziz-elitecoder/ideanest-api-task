// pkg/controllers/organization_controller.go
package controllers

import (
	"errors"

	"github.com/Abdulaziz-elitecoder/ideanest-api-task/pkg/database/mongodb/models"
	"github.com/Abdulaziz-elitecoder/ideanest-api-task/pkg/database/mongodb/repository"
)

// CreateOrganization creates a new organization.
func CreateOrganization(org models.Organization) (string, error) {
	// Call the repository to insert the organization into the database
	orgID, err := repository.InsertOrganization(org)
	if err != nil {
		return "", err
	}

	return orgID, nil
}

// ReadOrganization retrieves details of an organization.
func ReadOrganization(orgID string) (*models.Organization, error) {
	// Call the repository to get the organization from the database
	org, err := repository.GetOrganization(orgID)
	if err != nil {
		return nil, err
	}

	return org, nil
}

// UpdateOrganization updates details of an organization.
func UpdateOrganization(orgID string, updatedOrg models.Organization) error {
	// Call the repository to update the organization in the database
	err := repository.UpdateOrganization(orgID, updatedOrg)
	if err != nil {
		return err
	}

	return nil
}

// DeleteOrganization deletes an organization.
func DeleteOrganization(orgID string) error {
	// Call the repository to delete the organization from the database
	err := repository.DeleteOrganization(orgID)
	if err != nil {
		return err
	}

	return nil
}

// InviteUserToOrganization sends an invitation to a user to join an organization.
func InviteUserToOrganization(orgID, userEmail string) error {
	// TODO: Implement user invitation logic
	// This might involve sending an email to the user or updating the database with the invitation

	// For now, let's return an error to indicate that the functionality is not implemented
	return errors.New("InviteUserToOrganization not implemented")
}
