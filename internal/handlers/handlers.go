package handlers

import (
	"encoding/json"
	"net/http"

	driverSvc "github.com/bcgov/efv-icbc-api/internal/service/driver"
	insuranceSvc "github.com/bcgov/efv-icbc-api/internal/service/insurance"
	vehicleSvc "github.com/bcgov/efv-icbc-api/internal/service/vehicle"
)

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

// RegisterRoutes registers HTTP handlers on the provided mux.
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/driver/validate", func(w http.ResponseWriter, r *http.Request) {
		license := r.URL.Query().Get("license_number")
		province := r.URL.Query().Get("province")
		resp, status := driverSvc.DriverValidation(license, province)
		writeJSON(w, status, resp)
	})

	mux.HandleFunc("/api/v1/insurance/status", func(w http.ResponseWriter, r *http.Request) {
		vin := r.URL.Query().Get("vin")
		policy := r.URL.Query().Get("policy_number")
		resp, status := insuranceSvc.InsuranceStatus(vin, policy)
		writeJSON(w, status, resp)
	})

	mux.HandleFunc("/api/v1/vehicle/verify", func(w http.ResponseWriter, r *http.Request) {
		vin := r.URL.Query().Get("vin")
		plate := r.URL.Query().Get("plate")
		province := r.URL.Query().Get("province")
		resp, status := vehicleSvc.VehicleVerification(vin, plate, province)
		writeJSON(w, status, resp)
	})
}
