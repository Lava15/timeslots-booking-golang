package models

import "time"

type Booking struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	TimeslotID int       `json:"timeslot_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at,omitempty"`
}
