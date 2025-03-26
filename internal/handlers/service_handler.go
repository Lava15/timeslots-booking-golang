package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	repository "github.com/lava15/timeslots-booking-golang/internal/repositories"
	response "github.com/lava15/timeslots-booking-golang/pkg/responses"
)

func GetServices(w http.ResponseWriter, r *http.Request) {
	services, err := repository.GetAllServices()
	if err != nil {
		http.Error(w, "Failed to fetch services", http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, services)
}

func GetService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "Invalid service ID", http.StatusBadRequest)
		return
	}
	service, err := repository.GetServiceByID(id)
	if err != nil {
		http.Error(w, "Failed to fetch service", http.StatusInternalServerError)
		return
	}
	if service == nil {
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	}
	response.JSON(w, http.StatusOK, service)
}
