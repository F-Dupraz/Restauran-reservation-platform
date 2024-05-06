package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/segmentio/ksuid"

	"github.com/F-Dupraz/Restauran-reservation-platform.git/models"
	"github.com/F-Dupraz/Restauran-reservation-platform.git/repository"
	"github.com/F-Dupraz/Restauran-reservation-platform.git/server"
)

type InsterNewRestraurantRequest struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	City        string   `json:"city"`
	DaysOpen    []string `json:"days_open"`
	Capacity    []int    `json:"capacity"`
	Specialties []string `json:"specialties"`
}

type InsterNewRestraurantResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type GetAllRestaurantRequest struct {
	Offset int `json:"offset"`
}

type GetRestaurantByNameRequest struct {
	Offset int    `json:"offset"`
	Name   string `json:"name"`
}

type GetRestaurantByCityRequest struct {
	Offset int    `json:"offset"`
	City   string `json:"city"`
}

type GetRestaurantsResponse struct {
	Restaurants []models.Restaurant `json:"restaurants"`
}

type UpdateRestraurantRequest struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	DaysOpen    []string `json:"days_open"`
	Capacity    []int    `json:"capacity"`
	Specialties []string `json:"specialties"`
}

type UpdateRestraurantResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
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
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.UserToken{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*models.UserToken); ok && token.Valid {
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
				Owner:       claims.Id,
				DaysOpen:    request.DaysOpen,
				Capacity:    request.Capacity,
				Specialties: request.Specialties,
			}
			err = repository.InsterNewRestraurant(r.Context(), &restaurant)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(InsterNewRestraurantResponse{
				Success: true,
				Message: "Restaurant added successfully ;)",
			})
		}
	}
}

func GetAllRestaurants(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = GetAllRestaurantRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var offset = request.Offset
		restaurants, err := repository.GetAllRestaurants(r.Context(), offset)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
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
		var offset = request.Offset
		restaurants, err := repository.GetRestaurantByName(r.Context(), name, offset)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
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
		var offset = request.Offset
		restaurants, err := repository.GetRestaurantByCity(r.Context(), city, offset)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(GetRestaurantsResponse{
			Restaurants: restaurants,
		})

	}
}

func GetMyRestaurantsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.UserToken{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*models.UserToken); ok && token.Valid {
			restaurants, err := repository.GetMyRestaurants(r.Context(), claims.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(GetRestaurantsResponse{
				Restaurants: restaurants,
			})
		}
	}
}

func UpdateRestaurantHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.UserToken{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*models.UserToken); ok && token.Valid {
			var request = UpdateRestraurantRequest{}
			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			var updated_restaurant = models.Restaurant{
				Id:          request.Id,
				Name:        strings.ToLower(request.Name),
				DaysOpen:    request.DaysOpen,
				Capacity:    request.Capacity,
				Specialties: request.Specialties,
			}
			err = repository.UpdateRestaurant(r.Context(), &updated_restaurant, claims.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(UpdateRestraurantResponse{
				Success: true,
				Message: "Restaurant updated successfully ;)",
			})
		}
	}
}

func DeleteRestaurantHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.UserToken{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*models.UserToken); ok && token.Valid {
			var request = DeleteRestaurantRequest{}
			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			var id = request.Id
			var user_id = claims.Id
			err = repository.DeleteRestaurant(r.Context(), id, user_id)
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
}
