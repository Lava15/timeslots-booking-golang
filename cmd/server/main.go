package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lava15/timeslots-booking-golang/internal/db"
)

func main() {
	db.InitDB()
	r := mux.NewRouter()
	slog.Info("Starting server...")
	fmt.Println("Starting server...")
	http.ListenAndServe(":8080", r)
}
