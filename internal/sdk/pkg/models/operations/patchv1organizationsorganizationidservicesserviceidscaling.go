// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type PatchV1OrganizationsOrganizationIDServicesServiceIDScalingRequest struct {
	ServiceScalingPatchRequest *shared.ServiceScalingPatchRequest `request:"mediaType=application/json"`
	// ID of the organization that owns the service.
	OrganizationID string `pathParam:"style=simple,explode=false,name=organizationId"`
	// ID of the service to update scaling parameters.
	ServiceID string `pathParam:"style=simple,explode=false,name=serviceId"`
}

// PatchV1OrganizationsOrganizationIDServicesServiceIDScaling400ApplicationJSON - The server cannot or will not process the request due to something that is perceived to be a client error.
type PatchV1OrganizationsOrganizationIDServicesServiceIDScaling400ApplicationJSON struct {
	// Detailed error description.
	Error *string `json:"error,omitempty"`
	// HTTP status code.
	Status *float64 `json:"status,omitempty"`
}

// PatchV1OrganizationsOrganizationIDServicesServiceIDScaling200ApplicationJSON - Successful response
type PatchV1OrganizationsOrganizationIDServicesServiceIDScaling200ApplicationJSON struct {
	Result *shared.Service `json:"result,omitempty"`
	// HTTP status code.
	Status *float64 `json:"status,omitempty"`
}

type PatchV1OrganizationsOrganizationIDServicesServiceIDScalingResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	PatchV1OrganizationsOrganizationIDServicesServiceIDScaling200ApplicationJSONObject *PatchV1OrganizationsOrganizationIDServicesServiceIDScaling200ApplicationJSON
	// The server cannot or will not process the request due to something that is perceived to be a client error.
	PatchV1OrganizationsOrganizationIDServicesServiceIDScaling400ApplicationJSONObject *PatchV1OrganizationsOrganizationIDServicesServiceIDScaling400ApplicationJSON
}