package db

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func initDB() {
	connStr := ""
	var err error
	DB, err = sql.Open("pg", connStr)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}
	_, err = DB.Exec(
		`CREATE TABLE IF NOT EXISTS users (
				id SERIAL PRIMARY KEY,
				username TEXT NOT NULL,
				password TEXT NOT NULL
				is_admin BOOLEAN NOT NULL
		);
		CREATE TABLE IF NOT EXISTS slots (
			id SERIAL PRIMARY KEY,
			business_id INTEGER NOT NULL,
			start_time TIMESTAMP NOT NULL,
			end_time TIMESTAMP NOT NULL
			booked BOOLEAN NOT NULL
			booked_by INTEGER
		);
		`)
	if err != nil {
		log.Fatal("Error creating tables: ", err)
	}
}
