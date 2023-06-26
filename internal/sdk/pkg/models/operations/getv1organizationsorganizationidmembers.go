// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"clickhouse/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetV1OrganizationsOrganizationIDMembersRequest struct {
	// ID of the requested organization.
	OrganizationID string `pathParam:"style=simple,explode=false,name=organizationId"`
}

// GetV1OrganizationsOrganizationIDMembers400ApplicationJSON - The server cannot or will not process the request due to something that is perceived to be a client error.
type GetV1OrganizationsOrganizationIDMembers400ApplicationJSON struct {
	// Detailed error description.
	Error *string `json:"error,omitempty"`
	// HTTP status code.
	Status *float64 `json:"status,omitempty"`
}

// GetV1OrganizationsOrganizationIDMembers200ApplicationJSON - Successful response
type GetV1OrganizationsOrganizationIDMembers200ApplicationJSON struct {
	Result []shared.Member `json:"result,omitempty"`
	// HTTP status code.
	Status *float64 `json:"status,omitempty"`
}

type GetV1OrganizationsOrganizationIDMembersResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	GetV1OrganizationsOrganizationIDMembers200ApplicationJSONObject *GetV1OrganizationsOrganizationIDMembers200ApplicationJSON
	// The server cannot or will not process the request due to something that is perceived to be a client error.
	GetV1OrganizationsOrganizationIDMembers400ApplicationJSONObject *GetV1OrganizationsOrganizationIDMembers400ApplicationJSON
}
