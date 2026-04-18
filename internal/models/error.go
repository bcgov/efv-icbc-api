package models

type ErrorResponse struct {
	Code       string `json:"code"`
	Details    string `json:"details"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"http_status,omitempty"`
}
