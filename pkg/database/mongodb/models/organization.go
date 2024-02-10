package models

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UnmarshalOrganization(data []byte) (Organization, error) {
	var r Organization
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Organization) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Organization struct {
	Id                     primitive.ObjectID   `bson:"_id" json:"_id"`
	Name                   string               `bson:"name" form:"name" json:"name,omitempty"  validate:"required"`
	Description            string               `bson:"description" form:"description" json:"description,omitempty"  validate:"required"`
	OrganizationMembersIDs []primitive.ObjectID `bson:"organization_members" json:"organization_members,omitempty"`
}
