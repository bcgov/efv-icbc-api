package mocks

import "github.com/bcgov/efv-icbc-api/internal/models"

// NewVehicleInfo returns a mock VehicleInfo (no business validation).
func NewVehicleInfo(vin, plate, province string) models.VehicleInfo {
	prov := province
	if prov == "" {
		prov = "BC"
	}
	return models.VehicleInfo{
		VIN:                vin,
		PlateNumber:        plate,
		Province:           prov,
		Make:               "Toyota",
		Model:              "Corolla",
		Year:               2019,
		RegistrationStatus: "registered",
		RegistrationExpiry: "2025-06-30",
		InsuranceStatus:    "active",
		InsuranceExpiry:    "2024-06-30",
		IsCommercial:       false,
	}
}
