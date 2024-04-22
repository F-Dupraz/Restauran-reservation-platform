package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/F-Dupraz/Restauran-reservation-platform.git/handlers"
	"github.com/F-Dupraz/Restauran-reservation-platform.git/server"
)

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.InsterNewRestraurantHandler(s)).Methods(http.MethodPost)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v.\n", err)
	}

	PORT := os.Getenv("RESTAURANT_PORT")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port: PORT,
	})

	if err != nil {
		log.Fatalf("Error creating server %v.\n", err)
	}

	s.Start(BindRoutes)
}
