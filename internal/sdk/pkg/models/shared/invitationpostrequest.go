// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// InvitationPostRequestRole - Role of the member in the organization.
type InvitationPostRequestRole string

const (
	InvitationPostRequestRoleAdmin     InvitationPostRequestRole = "admin"
	InvitationPostRequestRoleDeveloper InvitationPostRequestRole = "developer"
)

func (e InvitationPostRequestRole) ToPointer() *InvitationPostRequestRole {
	return &e
}

func (e *InvitationPostRequestRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "admin":
		fallthrough
	case "developer":
		*e = InvitationPostRequestRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InvitationPostRequestRole: %v", v)
	}
}

type InvitationPostRequest struct {
	// Email of the invited user. Only a user with this email can join using the invitation. The email is stored in a lowercase form.
	Email string `json:"email"`
	// Role of the member in the organization.
	Role InvitationPostRequestRole `json:"role"`
}
