package repository

import (
	"github.com/google/uuid"
	"github.com/lava15/timeslots-booking-golang/internal/db"
	"github.com/lava15/timeslots-booking-golang/internal/models"
)

func GetAllServices() ([]models.Service, error) {
	rows, err := db.DB.Query(`SELECT * FROM services WHERE deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var s models.Service
		err := rows.Scan(&s.ID, &s.Name, &s.Description, &s.Slug, &s.Price, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
		if err != nil {
			return nil, err
		}
		services = append(services, s)
	}
	return services, nil
}

func GetServiceByID(id uuid.UUID) (models.Service, error) {
	var s models.Service
	err := db.DB.QueryRow(`SELECT * FROM services WHERE id = ? AND deleted_at IS NULL`, id).Scan(&s.ID, &s.Name, &s.Description, &s.Slug, &s.Price, &s.CreatedAt, &s.UpdatedAt, &s.DeletedAt)
	if err != nil {
		return models.Service{}, err
	}
	return s, nil
}
