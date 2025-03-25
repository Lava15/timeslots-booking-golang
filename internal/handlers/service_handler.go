package handlers

import (
	"encoding/json"
	"net/http"

	repository "github.com/lava15/timeslots-booking-golang/internal/repositories"
)

func GetServices(w http.ResponseWriter, r *http.Request) {
	services, err := repository.GetAllServices()
	if err != nil {
		http.Error(w, "Failed to fetch services", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetService(w http.ResponseWriter, r *http.Request) {

}
