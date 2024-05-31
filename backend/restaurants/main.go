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
	r.HandleFunc("/api/restaurants", handlers.InsterNewRestraurantHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/api/restaurants", handlers.DeleteRestaurantHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/api/restaurants", handlers.UpdateRestaurantHandler(s)).Methods(http.MethodPatch)
	r.HandleFunc("/api/restaurants", handlers.GetAllRestaurants(s)).Methods(http.MethodGet)
	r.HandleFunc("/api/restaurants/mines", handlers.GetMyRestaurants(s)).Methods(http.MethodGet)
	r.HandleFunc("/api/restaurants/{id}", handlers.GetRestaurantById(s)).Methods(http.MethodGet)
	r.HandleFunc("/api/restaurants/name", handlers.GetRestaurantByNameHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/api/restaurants/city", handlers.GetRestaurantByCityHandler(s)).Methods(http.MethodGet)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v.\n", err)
	}

	PORT := os.Getenv("RESTAURANT_PORT")
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
