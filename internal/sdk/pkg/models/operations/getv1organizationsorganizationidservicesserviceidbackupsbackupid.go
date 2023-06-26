// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"clickhouse/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetV1OrganizationsOrganizationIDServicesServiceIDBackupsBackupIDRequest struct {
	// ID of the requested backup.
	BackupID string `pathParam:"style=simple,explode=false,name=backupId"`
	// ID of the organization that owns the backup.
	OrganizationID string `pathParam:"style=simple,explode=false,name=organizationId"`
	// ID of the service the backup was created from.
	ServiceID string `pathParam:"style=simple,explode=false,name=serviceId"`
}

// GetV1OrganizationsOrganizationIDServicesServiceIDBackupsBackupID400ApplicationJSON - The server cannot or will not process the request due to something that is perceived to be a client error.
type GetV1OrganizationsOrganizationIDServicesServiceIDBackupsBackupID400ApplicationJSON struct {
	// Detailed error description.
	Error *string `json:"error,omitempty"`
	// HTTP status code.
	Status *float64 `json:"status,omitempty"`
}

// GetV1OrganizationsOrganizationIDServicesServiceIDBackupsBackupID200ApplicationJSON - Successful response
type GetV1OrganizationsOrganizationIDServicesServiceIDBackupsBackupID200ApplicationJSON struct {
	Result *shared.Backup `json:"result,omitempty"`
	// HTTP status code.
	Status *float64 `json:"status,omitempty"`
}

type GetV1OrganizationsOrganizationIDServicesServiceIDBackupsBackupIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	GetV1OrganizationsOrganizationIDServicesServiceIDBackupsBackupID200ApplicationJSONObject *GetV1OrganizationsOrganizationIDServicesServiceIDBackupsBackupID200ApplicationJSON
	// The server cannot or will not process the request due to something that is perceived to be a client error.
	GetV1OrganizationsOrganizationIDServicesServiceIDBackupsBackupID400ApplicationJSONObject *GetV1OrganizationsOrganizationIDServicesServiceIDBackupsBackupID400ApplicationJSON
}
