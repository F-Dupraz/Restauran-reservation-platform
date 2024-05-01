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
