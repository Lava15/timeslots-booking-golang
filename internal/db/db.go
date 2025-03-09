package db

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func InitDB() {
	connStr := "user=postgres dbname=booking password=postgres port=5432 sslmode=disable"
	var err error
	DB, err = sql.Open("pg", connStr)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
}

func CreateTables() {
	_, err := DB.Exec(`
			CREATE TABLE IF NOT EXISTS users (
				id SERIAL PRIMARY KEY,
				email TEXT UNIQUE NOT NULL,
				password TEXT NOT NULL
				is_admin BOOLEAN DEFAULT FALSE
			);
		`)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS services (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description TEXT,
			price DECIMAL(10, 2) NOT NULL
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP
		);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created successfully")
}
