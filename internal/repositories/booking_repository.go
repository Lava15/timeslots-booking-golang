package repository

import (
	"encoding/json"
	"net/http"

	"github.com/lava15/timeslots-booking-golang/internal/db"
	"github.com/lava15/timeslots-booking-golang/internal/models"
)

func GetAllBookings(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT user_id, timeslot_id, FROM bookings WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var booking models.Booking
		err := rows.Scan(&booking.UserID, &booking.TimeslotID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bookings = append(bookings, booking)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}
