package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	response := HealthCheckResponse{
		Message: "Yep! Good as new!",
	}
	json.NewEncoder(w).Encode(response)
	return
}
