package repository

import (
	"encoding/json"
	"net/http"

	"github.com/lava15/timeslots-booking-golang/internal/db"
	"github.com/lava15/timeslots-booking-golang/internal/models"
)

func GetAllServices(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`SELECT * FROM services WHERE deleted_at IS NULL`)
	if err != nil {
		http.Error(w, "Failed to fetch services", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var s models.Service
		err := rows.Scan(&s.ID, &s.Name, &s.Description, &s.Slug, &s.Price, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
		if err != nil {
			http.Error(w, "Failed to fetch services", http.StatusInternalServerError)
			return
		}
		services = append(services, s)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}
