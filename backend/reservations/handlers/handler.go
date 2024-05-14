package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

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
	DayInt       string `json:"day_int"`
	From         string `json:"from"`
	To           string `json:"to"`
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
	From      string `json:"from"`
	To        string `json:"to"`
	NumGuests int    `json:"num_guests"`
}

type UpdateReservationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type DeleteReservationRequest struct {
	Id string `json:"id"`
}

type DeleteReservationResponse struct {
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
	From         string `json:"from"`
	To           string `json:"to"`
	NumGuests    int    `json:"num_guests"`
	IsDone       bool   `json:"is_done"`
}

type GetReservationByDayRequest struct {
	Day          string `json:"day"`
	RestaurantId string `json:"restaurant_id"`
}

type GetReservationByRidRequest struct {
	RestaurantId string `json:"restaurant_id"`
}

type GetReservationByDayResponse struct {
	Reservations []models.Reservation
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
			if request.RestaurantId == "" || request.UserId == "" || request.NumGuests == 0 || request.NumGuests < 0 {
				http.Error(w, "Bad request. Maybe you forgot an argument.", http.StatusBadRequest)
				return
			}
			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			DayParsed, _ := time.Parse(time.DateOnly, request.Day)
			fmt.Println(request.Day)
			DayOfTheWeek := DayParsed.Weekday()
			fmt.Println(DayOfTheWeek)
			fmt.Println(int(DayOfTheWeek))
			var DayOfTheWeekArr = []int{int(DayOfTheWeek)}

			var reservation = models.Reservation{
				Id:           id.String(),
				UserId:       claims.Id,
				RestaurantId: request.RestaurantId,
				Day:          request.Day,
				DayInt:       DayOfTheWeekArr,
				From:         request.From,
				To:           request.To,
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
			DayParsed, _ := time.Parse(time.DateOnly, request.Day)
			DayOfTheWeek := DayParsed.Weekday()
			fmt.Println(DayOfTheWeek)
			fmt.Println(int(DayOfTheWeek))
			var DayOfTheWeekArr = []int{int(DayOfTheWeek)}

			var reservation = models.Reservation{
				Id:        request.Id,
				Day:       request.Day,
				DayInt:    DayOfTheWeekArr,
				From:      request.From,
				To:        request.To,
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

func DeleteReservationHandler(s server.Server) http.HandlerFunc {
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
			var request = DeleteReservationRequest{}
			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if request.Id == "" {
				http.Error(w, "Bad request. Maybe you forgot the id.", http.StatusBadRequest)
				return
			}
			err = repository.DeleteReservation(r.Context(), claims.Id, request.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(DeleteReservationResponse{
				Success: true,
				Message: "Reservation deleted correctly",
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
		if _, ok := token.Claims.(*models.UserToken); ok && token.Valid {
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
			reservation, err := repository.GetReservationById(r.Context(), request.Id)
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
				From:         reservation.From,
				To:           reservation.To,
				NumGuests:    reservation.NumGuests,
				IsDone:       reservation.IsDone,
			})
		}
	}
}

func GetReservationByDayHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.UserToken{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if _, ok := token.Claims.(*models.UserToken); ok && token.Valid {
			var request = GetReservationByDayRequest{}
			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if request.RestaurantId == "" || request.Day[0] == 0 {
				http.Error(w, "Bad request. Maybe you forgot something.", http.StatusBadRequest)
				return
			}
			DayParsed, _ := time.Parse(time.DateOnly, request.Day)
			DayOfTheWeek := DayParsed.Weekday()
			fmt.Println(DayOfTheWeek)
			fmt.Println(int(DayOfTheWeek))
			var DayOfTheWeekArr = [1]int{int(DayOfTheWeek)}
			reservations, err := repository.GetReservationsByDay(r.Context(), DayOfTheWeekArr, request.RestaurantId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(GetReservationByDayResponse{
				Reservations: reservations,
			})
		}
	}
}

func GetReservationsByRidHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		token, err := jwt.ParseWithClaims(tokenString, &models.UserToken{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if _, ok := token.Claims.(*models.UserToken); ok && token.Valid {
			var request = GetReservationByRidRequest{}
			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if request.RestaurantId == "" {
				http.Error(w, "Bad request. Maybe you forgot something.", http.StatusBadRequest)
				return
			}
			reservations, err := repository.GetReservationsByRid(r.Context(), request.RestaurantId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(GetReservationByDayResponse{
				Reservations: reservations,
			})
		}
	}
}
