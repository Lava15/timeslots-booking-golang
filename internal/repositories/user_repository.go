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

func (r *UserRepository) CreateUser(u models.User) error {
	var deletedAt interface{}
	if u.DeletedAt == "" {
		deletedAt = nil
	} else {
		deletedAt = u.DeletedAt
	}
	_, err := r.db.Exec(`
		INSERT INTO user (id, email, password, is_admin, created_at, updated_at, deleted_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		u.ID, u.Email, u.Password, u.IsAdmin, u.CreatedAt, u.UpdatedAt, deletedAt)

	return err
}

func (r *UserRepository) UserExistsByEmail(email string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(`
	SELECT EXISTS (SELECT 1 FROM users WHERE email = ? AND deleted_at IS NULL)`, email).Scan(&exists)
	return exists, err
}

func (r *UserRepository) HasUsers() (bool, error) {
	var exists bool
	err := r.db.QueryRow(`
		SELECT EXISTS (SELECT 1 FROM users WHERE deleted_at IS NULL)`).Scan(&exists)
	if err != nil {
		return false, nil
	}
	return exists, nil
}
