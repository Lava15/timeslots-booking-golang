package utils

import (
	"time"

	"github.com/google/uuid"
	"github.com/lava15/timeslots-booking-golang/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func GenerateFakeUsers(n int) []models.User {
	users := make([]models.User, n)
	existingEmails := make(map[string]bool)
	for i := 0; i < n; i++ {
		var email string
		for {
			email = gofakeit.Email()
			if !existingEmails[email] {
				existingEmails[email] = true
				break
			}
		}
		password := gofakeit.Password(true, true, true, true, false, 12)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		now := time.Now().Format(time.RFC3339)
		users[i] = models.User{
			ID:        uuid.New(),
			Email:     email,
			Password:  string(hashedPassword),
			IsAdmin:   gofakeit.Bool(),
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: "",
		}
	}
	return users
}
