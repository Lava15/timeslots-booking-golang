package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "postgres://postgres:postgres@db:5432/booking?sslmode=disable"
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
