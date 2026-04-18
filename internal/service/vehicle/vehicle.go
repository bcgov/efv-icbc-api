package vehicle

import (
	"fmt"
	"net/http"

	"github.com/bcgov/efv-icbc-api/internal/mocks"
	"github.com/bcgov/efv-icbc-api/internal/models"
)

// VehicleVerification enforces business rules for vehicle entity.
func VehicleVerification(vin, plate, province string) (models.VehicleVerificationResponse, int) {
	prov := province
	if prov == "" {
		prov = "BC"
	}
	if prov != "BC" {
		return models.VehicleVerificationResponse{
			Success: false,
			Error: &models.ErrorResponse{
				Code:       "unsupported_province",
				Message:    "Province not supported",
				Details:    fmt.Sprintf("Only 'BC' is supported; got '%s'", prov),
				HTTPStatus: http.StatusBadRequest,
			},
		}, http.StatusBadRequest
	}

	info := mocks.NewVehicleInfo(vin, plate, prov)
	return models.VehicleVerificationResponse{Success: true, Data: &info}, http.StatusOK
}
