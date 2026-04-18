package mocks

import "github.com/bcgov/efv-icbc-api/internal/models"

// NewDriverInfo returns a mock DriverInfo (no business validation).
func NewDriverInfo(license, province string) models.DriverInfo {
	prov := province
	if prov == "" {
		prov = "BC"
	}
	return models.DriverInfo{
		LicenseNumber: license,
		Province:      prov,
		Status:        "valid",
		Class:         "5",
		IsProvisional: false,
		IsSuspended:   false,
		IssueDate:     "2018-05-01",
		ExpiryDate:    "2028-05-01",
		Restrictions:  []string{},
	}
}
