// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"clickhouse/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetV1OrganizationsOrganizationIDKeysKeyIDRequest struct {
	// ID of the requested key.
	KeyID string `pathParam:"style=simple,explode=false,name=keyId"`
	// ID of the requested organization.
	OrganizationID string `pathParam:"style=simple,explode=false,name=organizationId"`
}

// GetV1OrganizationsOrganizationIDKeysKeyID400ApplicationJSON - The server cannot or will not process the request due to something that is perceived to be a client error.
type GetV1OrganizationsOrganizationIDKeysKeyID400ApplicationJSON struct {
	// Detailed error description.
	Error *string `json:"error,omitempty"`
	// HTTP status code.
	Status *float64 `json:"status,omitempty"`
}

// GetV1OrganizationsOrganizationIDKeysKeyID200ApplicationJSON - Successful response
type GetV1OrganizationsOrganizationIDKeysKeyID200ApplicationJSON struct {
	Result *shared.APIKey `json:"result,omitempty"`
	// HTTP status code.
	Status *float64 `json:"status,omitempty"`
}

type GetV1OrganizationsOrganizationIDKeysKeyIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	GetV1OrganizationsOrganizationIDKeysKeyID200ApplicationJSONObject *GetV1OrganizationsOrganizationIDKeysKeyID200ApplicationJSON
	// The server cannot or will not process the request due to something that is perceived to be a client error.
	GetV1OrganizationsOrganizationIDKeysKeyID400ApplicationJSONObject *GetV1OrganizationsOrganizationIDKeysKeyID400ApplicationJSON
}
