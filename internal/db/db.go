package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
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
