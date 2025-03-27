package db

import (
	"database/sql"
	"fmt"

	"github.com/lava15/timeslots-booking-golang/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
	)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Error opening database: %q", err)
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		fmt.Printf("Error pinging database: %q", err)
		panic(err)
	}
	fmt.Println("Connected to database")
}
