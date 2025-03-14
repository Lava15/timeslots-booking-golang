package models

import (
	"time"

	"github.com/google/uuid"
)

type Slot struct {
	ID         uuid.UUID `json:"id"`
	BusinessID uuid.UUID `json:"business_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	Booked     bool      `json:"is_booked"`
	BookedBy   *int      `json:"booked_by"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateAt   time.Time `json:"updated_at"`
	DeleteAt   time.Time `json:"deleted_at,omitempty"`
}
