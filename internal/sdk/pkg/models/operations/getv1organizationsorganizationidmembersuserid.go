// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetV1OrganizationsOrganizationIDMembersUserIDRequest struct {
	// ID of the organization the member is part of.
	OrganizationID string `pathParam:"style=simple,explode=false,name=organizationId"`
	// ID of the requested user.
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

// GetV1OrganizationsOrganizationIDMembersUserID400ApplicationJSON - The server cannot or will not process the request due to something that is perceived to be a client error.
type GetV1OrganizationsOrganizationIDMembersUserID400ApplicationJSON struct {
	// Detailed error description.
	Error *string `json:"error,omitempty"`
	// HTTP status code.
	Status *float64 `json:"status,omitempty"`
}

// GetV1OrganizationsOrganizationIDMembersUserID200ApplicationJSON - Successful response
type GetV1OrganizationsOrganizationIDMembersUserID200ApplicationJSON struct {
	Result *shared.Member `json:"result,omitempty"`
	// HTTP status code.
	Status *float64 `json:"status,omitempty"`
}

type GetV1OrganizationsOrganizationIDMembersUserIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	GetV1OrganizationsOrganizationIDMembersUserID200ApplicationJSONObject *GetV1OrganizationsOrganizationIDMembersUserID200ApplicationJSON
	// The server cannot or will not process the request due to something that is perceived to be a client error.
	GetV1OrganizationsOrganizationIDMembersUserID400ApplicationJSONObject *GetV1OrganizationsOrganizationIDMembersUserID400ApplicationJSON
}