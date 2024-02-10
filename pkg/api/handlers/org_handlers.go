// pkg/api/handlers/org_handlers.go
package handlers

import (
	"net/http"

	"github.com/Abdulaziz-elitecoder/ideanest-api-task/pkg/controllers"
	"github.com/Abdulaziz-elitecoder/ideanest-api-task/pkg/database/mongodb/models"
	"github.com/gin-gonic/gin"
)

// CreateOrganization handles the creation of a new organization.
func CreateOrganization(c *gin.Context) {
	var org models.Organization
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the controller to handle organization creation
	orgID, err := controllers.CreateOrganization(org)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create organization"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"organization_id": orgID})
}

// ReadOrganization handles the retrieval of organization details.
func ReadOrganization(c *gin.Context) {
	orgID := c.Param("organization_id")

	// Call the controller to handle organization retrieval
	org, err := controllers.ReadOrganization(orgID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}

	c.JSON(http.StatusOK, org)
}

// UpdateOrganization handles the update of organization details.
func UpdateOrganization(c *gin.Context) {
	orgID := c.Param("organization_id")

	var updatedOrg models.Organization
	if err := c.ShouldBindJSON(&updatedOrg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the controller to handle organization update
	err := controllers.UpdateOrganization(orgID, updatedOrg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update organization"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organization updated successfully"})
}

// DeleteOrganization handles the deletion of an organization.
func DeleteOrganization(c *gin.Context) {
	orgID := c.Param("organization_id")

	// Call the controller to handle organization deletion
	err := controllers.DeleteOrganization(orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete organization"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organization deleted successfully"})
}

// InviteUserToOrganization handles the invitation of a user to an organization.
func InviteUserToOrganization(c *gin.Context) {
	orgID := c.Param("organization_id")

	var inviteReq models.InviteRequest
	if err := c.ShouldBindJSON(&inviteReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the controller to handle user invitation
	err := controllers.InviteUserToOrganization(orgID, inviteReq.UserEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to invite user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User invited to organization successfully"})
}
