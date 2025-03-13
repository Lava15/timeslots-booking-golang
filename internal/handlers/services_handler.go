package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lava15/timeslots-booking-golang/internal/db"
	"github.com/lava15/timeslots-booking-golang/internal/models"
)

func GetServices(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`
    SELECT * FROM services WHERE deleted_at IS NULL
  `)
	if err != nil {
		http.Error(w, "Failed to fetch services", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var s models.Service
		err := rows.Scan(&s.ID, &s.Name, &s.Price, &s.Duration, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
		if err != nil {
			http.Error(w, "Failed to fetch services", http.StatusInternalServerError)
			return
		}
		services = append(services, s)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func GetService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid service ID", http.StatusBadRequest)
		return
	}

	var s models.Service
	err = db.DB.QueryRow(`
      SELECT id, name, description, slug, price, created_at, updated_at, deleted_at
      FROM services
      WHERE id = $1 AND deleted_at IS NULL
  `, id).Scan(&s.ID, &s.Name, &s.Description, &s.Slug, &s.Price, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
	if err == sql.ErrNoRows {
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Failed to fetch service", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}
