package insurance

import (
	"net/http"

	"github.com/bcgov/efv-icbc-api/internal/mocks"
	"github.com/bcgov/efv-icbc-api/internal/models"
)

// InsuranceStatus returns insurance status (no province rules applied).
func InsuranceStatus(vin, policy string) (models.InsuranceStatusResponse, int) {
	info := mocks.NewInsuranceInfo(vin, policy)
	return models.InsuranceStatusResponse{Success: true, Data: &info}, http.StatusOK
}
