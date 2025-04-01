package repository_test

import (
	"database/sql"
	"testing"
)

func TestSeedUsers(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	userRepo := NewUserRepository(db)

	// Seed users
	err = userRepo.SeedUsers()
	if err != nil {
		t.Fatalf("Failed to seed users: %v", err)
	}

	// Verify seeding
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to count users: %v", err)
	}

	if count == 0 {
		t.Fatal("No users seeded")
	}
}
