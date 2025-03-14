package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lava15/timeslots-booking-golang/internal/db"
)

func main() {

	if err := godotenv.Load(); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			slog.Info("No .env file found")
		} else {
			slog.Error("Error loading .env file", err)
		}
	}
	PORT := os.Getenv("PORT")

	db.InitDB()
	r := mux.NewRouter()
	slog.Info("Starting server...")
	fmt.Println("Starting server...")
	slog.Info("Starting server on port", "port", PORT)
	if err := http.ListenAndServe(":"+PORT, r); err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}
