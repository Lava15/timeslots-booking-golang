package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	,"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/lava15/timeslots-booking-golang/config"
	"github.com/lava15/timeslots-booking-golang/internal/db"
)

var wg sync.WaitGroup

func main() {
	config.Init()
	r := mux.NewRouter()
	PORT := os.Getenv("PORT")
	srv := http.Server{
		Addr:    ":" + PORT,
		Handler: r,
	}
	db.InitDB()

	go func() {
		slog.Info("Starting server on port", "port", PORT)
		if err := srv.ListenAndServe(); err != nil {
			slog.Error("Failed to start server", "error", err)
		}
	}()
	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-shutdown.Done()
	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", "error", err)
	}
	fmt.Println("Server shutdown successfully")
	fmt.Println("waiting for all goroutines to finish")
	wg.Wait()
	fmt.Println("All goroutines finished")
}
