package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/segmentio/ksuid"

	"github.com/F-Dupraz/Restauran-reservation-platform.git/models"
	"github.com/F-Dupraz/Restauran-reservation-platform.git/repository"
	"github.com/F-Dupraz/Restauran-reservation-platform.git/server"
)

type InsterNewRestraurantRequest struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	City        string   `json:"city"`
	Owner       string   `json:"owner"`
	DaysOpen    []string `json:"days_open"`
	Specialties []string `json:"specialties"`
}

type InsterNewRestraurantResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type GetRestaurantByNameRequest struct {
	Name string `json:"name"`
}

type GetRestaurantByCityRequest struct {
	City string `json:"city"`
}

type GetRestaurantsResponse struct {
	Restaurants []models.Restaurant `json:"restaurants"`
}

type DeleteRestaurantRequest struct {
	Id string `json:"id"`
}

type DeleteRestaurantResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func InsterNewRestraurantHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = InsterNewRestraurantRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var restaurant = models.Restaurant{
			Id:          id.String(),
			Name:        strings.ToLower(request.Name),
			City:        strings.ToLower(request.City),
			Owner:       strings.ToLower(request.Owner),
			DaysOpen:    request.DaysOpen,
			Specialties: request.Specialties,
		}
		err = repository.InsterNewRestraurant(r.Context(), &restaurant)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(InsterNewRestraurantResponse{
			Success: true,
			Message: "Restaurant added successfully ;)",
		})

	}
}

func GetAllRestaurants(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restaurants, err := repository.GetAllRestaurants(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(GetRestaurantsResponse{
			Restaurants: restaurants,
		})
	}
}

func GetRestaurantByNameHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = GetRestaurantByNameRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var name = request.Name
		restaurants, err := repository.GetRestaurantByName(r.Context(), name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(GetRestaurantsResponse{
			Restaurants: restaurants,
		})

	}
}

func GetRestaurantByCityHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = GetRestaurantByCityRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var city = request.City
		restaurants, err := repository.GetRestaurantByCity(r.Context(), city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(GetRestaurantsResponse{
			Restaurants: restaurants,
		})

	}
}

func DeleteRestaurantHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = DeleteRestaurantRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var id = request.Id
		err = repository.DeleteRestaurant(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(DeleteRestaurantResponse{
			Success: true,
			Message: "Restaurant deleted successfully ;)",
		})
	}
}
