// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"clickhouse/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetV1OrganizationsOrganizationIDKeysRequest struct {
	// ID of the requested organization.
	OrganizationID string `pathParam:"style=simple,explode=false,name=organizationId"`
}

// GetV1OrganizationsOrganizationIDKeys400ApplicationJSON - The server cannot or will not process the request due to something that is perceived to be a client error.
type GetV1OrganizationsOrganizationIDKeys400ApplicationJSON struct {
	// Detailed error description.
	Error *string `json:"error,omitempty"`
}

// GetV1OrganizationsOrganizationIDKeys200ApplicationJSON - Successful response
type GetV1OrganizationsOrganizationIDKeys200ApplicationJSON struct {
	Result []shared.APIKey `json:"result,omitempty"`
}

type GetV1OrganizationsOrganizationIDKeysResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	GetV1OrganizationsOrganizationIDKeys200ApplicationJSONObject *GetV1OrganizationsOrganizationIDKeys200ApplicationJSON
	// The server cannot or will not process the request due to something that is perceived to be a client error.
	GetV1OrganizationsOrganizationIDKeys400ApplicationJSONObject *GetV1OrganizationsOrganizationIDKeys400ApplicationJSON
}
