package models

type VehicleInfo struct {
	InsuranceExpiry    string `json:"insurance_expiry"`
	InsuranceStatus    string `json:"insurance_status"`
	IsCommercial       bool   `json:"is_commercial"`
	Make               string `json:"make"`
	Model              string `json:"model"`
	PlateNumber        string `json:"plate_number"`
	Province           string `json:"province"`
	RegistrationExpiry string `json:"registration_expiry"`
	RegistrationStatus string `json:"registration_status"`
	VIN                string `json:"vin"`
	Year               int    `json:"year"`
}

type VehicleVerificationResponse struct {
	Data    *VehicleInfo   `json:"data,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
	Success bool           `json:"success"`
}
