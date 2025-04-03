package main

import (
	"fmt"
	"os"

	"github.com/lava15/timeslots-booking-golang/config"
	"github.com/lava15/timeslots-booking-golang/internal/db"
	repository "github.com/lava15/timeslots-booking-golang/internal/repositories"
	utils "github.com/lava15/timeslots-booking-golang/utils/seeds"
)

func main() {
	config.Init()
	if os.Getenv("ENV") == "production" {
		panic("seeds not allowed in production")
	}
	db.InitDB()

	userRepo := repository.NewUserRepository()
	hasUsers, err := userRepo.HasUsers()
	if err != nil {
		panic(fmt.Sprintf("failed to check if users exist: %v", err))
	}
	if !hasUsers {
		fakeUsers := utils.GenerateFakeUsers(10)
		for _, user := range fakeUsers {
			if err := userRepo.CreateUser(user); err != nil {
				panic(fmt.Sprintf("failed to create user: %v", err))
			}
		}
		fmt.Println("Fake users created successfully")
	} else {
		fmt.Println("Users already exist, skipping user creation")
	}
}
