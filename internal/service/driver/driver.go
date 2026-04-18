package driver

import (
	"fmt"
	"net/http"

	"github.com/bcgov/efv-icbc-api/internal/mocks"
	"github.com/bcgov/efv-icbc-api/internal/models"
)

// DriverValidation enforces business rules for driver entity.
func DriverValidation(license, province string) (models.DriverValidationResponse, int) {
	prov := province
	if prov == "" {
		prov = "BC"
	}
	if prov != "BC" {
		return models.DriverValidationResponse{
			Success: false,
			Error: &models.ErrorResponse{
				Code:       "unsupported_province",
				Message:    "Province not supported",
				Details:    fmt.Sprintf("Only 'BC' is supported; got '%s'", prov),
				HTTPStatus: http.StatusBadRequest,
			},
		}, http.StatusBadRequest
	}

	info := mocks.NewDriverInfo(license, prov)
	return models.DriverValidationResponse{Success: true, Data: &info}, http.StatusOK
}
