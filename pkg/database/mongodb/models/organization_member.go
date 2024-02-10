package models

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UnmarshalOrganizationMember(data []byte) (OrganizationMember, error) {
	var r OrganizationMember
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OrganizationMember) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type OrganizationMember struct {
	Id          primitive.ObjectID   `bson:"_id" json:"_id,omitempty"`
	Name        string               `form:"name,omitempty" json:"name,omitempty" validate:"required"`
	Email       string               `form:"email,omitempty" json:"email,omitempty" validate:"required"`
	Password    string               `form:"password,omitempty" json:"password,omitempty" validate:"required"`
	AccessLevel string               `form:"access_level" json:"access_level" bson:"access_level"`
	Invites     []primitive.ObjectID `json:"invites,omitempty"`
}
