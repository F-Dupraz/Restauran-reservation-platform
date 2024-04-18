package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/F-Dupraz/Restauran-reservation-platform.git/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status string `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HomeResponse {
			Message: "Everything is working just fine :)",
			Status: true,
		})
	}
}
