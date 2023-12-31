// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// InvitationRole - Role of the member in the organization.
type InvitationRole string

const (
	InvitationRoleAdmin     InvitationRole = "admin"
	InvitationRoleDeveloper InvitationRole = "developer"
)

func (e InvitationRole) ToPointer() *InvitationRole {
	return &e
}

func (e *InvitationRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "admin":
		fallthrough
	case "developer":
		*e = InvitationRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InvitationRole: %v", v)
	}
}

type Invitation struct {
	// Invitation creation timestamp. ISO-8601.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Email of the invited user. Only a user with this email can join using the invitation. The email is stored in a lowercase form.
	Email *string `json:"email,omitempty"`
	// Timestamp the invitation expires. ISO-8601.
	ExpireAt *time.Time `json:"expireAt,omitempty"`
	// Unique invitation ID.
	ID *string `json:"id,omitempty"`
	// Role of the member in the organization.
	Role *InvitationRole `json:"role,omitempty"`
}
