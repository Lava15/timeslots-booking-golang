package models

import "time"

type slot struct {
	ID         int       `json:"id"`
	BusinessID int       `json:"business_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	Booked     bool      `json:is_booked`
	BookedBy   *int      `json:booked_by`
	CreatedAt  time.Time `json:"created_at"`
}
