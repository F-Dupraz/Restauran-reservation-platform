package repository

import (
	"context"

	"github.com/F-Dupraz/Restauran-reservation-platform/reservations.git/models"
)

type ReservationRepository interface {
	CreateNewReservation(ctx context.Context, reservation *models.Reservation) error
	UpdateReservation(ctx context.Context, reservation *models.Reservation, id string) error
	GetReservationById(ctx context.Context, id string, user_id string) (*models.Reservation, error)
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

func GetReservationById(ctx context.Context, id string, user_id string) (*models.Reservation, error) {
	return reservationImplementation.GetReservationById(ctx, id, user_id)
}
