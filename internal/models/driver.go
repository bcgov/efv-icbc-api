package models

type DriverInfo struct {
	Class         string   `json:"class"`
	ExpiryDate    string   `json:"expiry_date"`
	IsProvisional bool     `json:"is_provisional"`
	IsSuspended   bool     `json:"is_suspended"`
	IssueDate     string   `json:"issue_date"`
	LicenseNumber string   `json:"license_number"`
	Province      string   `json:"province"`
	Restrictions  []string `json:"restrictions"`
	Status        string   `json:"status"`
}

type DriverValidationResponse struct {
	Data    *DriverInfo    `json:"data,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
	Success bool           `json:"success"`
}
