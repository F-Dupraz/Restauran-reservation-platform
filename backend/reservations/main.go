package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/F-Dupraz/Restauran-reservation-platform/reservations.git/handlers"
	"github.com/F-Dupraz/Restauran-reservation-platform/reservations.git/server"
)

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/api/reservations", handlers.CreateNewReservationHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/api/reservations", handlers.UpdateReservationHandler(s)).Methods(http.MethodPatch)
	r.HandleFunc("/api/reservations", handlers.DeleteReservationHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/api/my-reservations", handlers.GetMyReservations(s)).Methods(http.MethodGet)
	r.HandleFunc("/api/reservations/{id}", handlers.GetReservationByIdHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/api/reservations/day", handlers.GetReservationByDayHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/api/reservations/restaurant-id", handlers.GetReservationsByRidHandler(s)).Methods(http.MethodGet)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v.\n", err)
	}

	PORT := os.Getenv("RESERVATION_PORT")
	DATABASE_URL := os.Getenv("DATABASE_URL")
	JWT_TOKEN := os.Getenv("JWT_TOKEN")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_TOKEN,
		DatabaseURL: DATABASE_URL,
	})

	if err != nil {
		log.Fatalf("Error creating server %v.\n", err)
	}

	s.Start(BindRoutes)
}
