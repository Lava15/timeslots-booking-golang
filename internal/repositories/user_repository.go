package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/lava15/timeslots-booking-golang/internal/db"
	"github.com/lava15/timeslots-booking-golang/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: db.DB}
}

func (r *UserRepository) GetUserByID(id uuid.UUID) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow(`
		SELECT id, email, is_admin, created_at, updated_at, deleted_at 
		FROM users WHERE id = ? AND deleted_at IS NULL `, id).Scan(
		&u.ID, &u.Email, &u.IsAdmin, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt,
	)
	if err == sql.ErrNoRows {
		return &models.User{}, nil
	}
	if err != nil {
		return &models.User{}, err
	}
	return &u, nil
}
