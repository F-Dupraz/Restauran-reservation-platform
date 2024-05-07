package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/segmentio/ksuid"

	"github.com/F-Dupraz/Restauran-reservation-platform/reservations.git/models"
	"github.com/F-Dupraz/Restauran-reservation-platform/reservations.git/repository"
	"github.com/F-Dupraz/Restauran-reservation-platform/reservations.git/server"
)

type CreateNewReservationRequest struct {
	Id           string `json:"id"`
	UserId       string `json:"user_id"`
	RestaurantId string `json:"restaurant_id"`
	Day          string `json:"day"`
	NumGuests    int    `json:"num_guests"`
}

type CreateNewReservationResponse struct {
	Id      string `json:"id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type UpdateReservationRequest struct {
	Id        string `json:"id"`
	Day       string `json:"day"`
	NumGuests int    `json:"num_guests"`
}

type UpdateReservationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type GetReservationByIdRequest struct {
	Id string `json:"id"`
}

type GetReservationByIdResponse struct {
	Id           string `json:"id"`
	UserId       string `json:"user_id"`
	RestaurantId string `json:"restaurant_id"`
	Day          string `json:"day"`
	NumGuests    int    `json:"num_guests"`
	IsDone       bool   `json:"is_done"`
}

func CreateNewReservationHandler(s server.Server) http.HandlerFunc {
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
			var request = CreateNewReservationRequest{}
			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if request.Day == "" || request.RestaurantId == "" || request.UserId == "" || request.NumGuests == 0 || request.NumGuests < 0 {
				http.Error(w, "Bad request. Maybe you forgot an argument.", http.StatusBadRequest)
				return
			}
			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			var reservation = models.Reservation{
				Id:           id.String(),
				UserId:       claims.Id,
				RestaurantId: request.RestaurantId,
				Day:          request.Day,
				NumGuests:    request.NumGuests,
			}
			err = repository.CreateNewReservation(r.Context(), &reservation)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(CreateNewReservationResponse{
				Id:      reservation.Id,
				Success: true,
				Message: "Reservations created successfully ;)",
			})
		}
	}
}

func UpdateReservationHandler(s server.Server) http.HandlerFunc {
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
			var request = UpdateReservationRequest{}
			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if request.Day == "" || request.NumGuests == 0 || request.NumGuests < 0 || request.Id == "" {
				http.Error(w, "Bad request. Maybe you forgot an argument.", http.StatusBadRequest)
				return
			}
			var reservation = models.Reservation{
				Id:        request.Id,
				Day:       request.Day,
				NumGuests: request.NumGuests,
			}
			err = repository.UpdateReservation(r.Context(), &reservation, claims.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(UpdateReservationResponse{
				Success: true,
				Message: "Reservations updated successfully ;)",
			})
		}
	}
}

func GetReservationByIdHandler(s server.Server) http.HandlerFunc {
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
			var request = GetReservationByIdRequest{}
			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if request.Id == "" {
				http.Error(w, "Bad request. Maybe you forgot the id.", http.StatusBadRequest)
				return
			}
			reservation, err := repository.GetReservationById(r.Context(), request.Id, claims.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(GetReservationByIdResponse{
				Id:           reservation.Id,
				UserId:       reservation.UserId,
				RestaurantId: reservation.RestaurantId,
				Day:          reservation.Day,
				NumGuests:    reservation.NumGuests,
				IsDone:       reservation.IsDone,
			})
		}
	}
}
