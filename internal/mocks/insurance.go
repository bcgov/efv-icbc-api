package mocks

import "github.com/bcgov/efv-icbc-api/internal/models"

// NewInsuranceInfo returns a mock InsuranceInfo.
func NewInsuranceInfo(vin, policy string) models.InsuranceInfo {
	return models.InsuranceInfo{
		VIN:           vin,
		PolicyNumber:  policy,
		Status:        "active",
		CoverageType:  "comprehensive",
		EffectiveDate: "2023-01-01",
		ExpiryDate:    "2024-01-01",
		IsCancelled:   false,
	}
}
