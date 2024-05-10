package repository

import (
	"context"

	"github.com/F-Dupraz/Restauran-reservation-platform/reservations.git/models"
)

type ReservationRepository interface {
	CreateNewReservation(ctx context.Context, reservation *models.Reservation) error
	UpdateReservation(ctx context.Context, reservation *models.Reservation, id string) error
	DeleteReservation(ctx context.Context, user_id string, id string) error
	GetReservationById(ctx context.Context, id string) (*models.Reservation, error)
	GetReservationsByRid(ctx context.Context, restaurant_id string) ([]models.Reservation, error)
	GetReservationsByDay(ctx context.Context, day [1]int, restaurant_id string) ([]models.Reservation, error)
}

var reservationImplementation ReservationRepository

func SetRepository(repository ReservationRepository) {
	reservationImplementation = repository
}

func CreateNewReservation(ctx context.Context, reservation *models.Reservation) error {
	return reservationImplementation.CreateNewReservation(ctx, reservation)
}

func UpdateReservation(ctx context.Context, reservation *models.Reservation, id string) error {
	return reservationImplementation.UpdateReservation(ctx, reservation, id)
}

func DeleteReservation(ctx context.Context, user_id string, id string) error {
	return reservationImplementation.DeleteReservation(ctx, user_id, id)
}

func GetReservationById(ctx context.Context, id string) (*models.Reservation, error) {
	return reservationImplementation.GetReservationById(ctx, id)
}

func GetReservationsByRid(ctx context.Context, restaurant_id string) ([]models.Reservation, error) {
	return reservationImplementation.GetReservationsByRid(ctx, restaurant_id)
}

func GetReservationsByDay(ctx context.Context, day [1]int, restaurant_id string) ([]models.Reservation, error) {
	return reservationImplementation.GetReservationsByDay(ctx, day, restaurant_id)
}
