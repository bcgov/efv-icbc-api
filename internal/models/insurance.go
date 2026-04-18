package models

type InsuranceInfo struct {
	CancellationDate string `json:"cancellation_date"`
	CoverageType     string `json:"coverage_type"`
	EffectiveDate    string `json:"effective_date"`
	ExpiryDate       string `json:"expiry_date"`
	IsCancelled      bool   `json:"is_cancelled"`
	PolicyNumber     string `json:"policy_number"`
	Status           string `json:"status"`
	VIN              string `json:"vin"`
}

type InsuranceStatusResponse struct {
	Data    *InsuranceInfo `json:"data,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
	Success bool           `json:"success"`
}
