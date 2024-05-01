package repository

import (
	"context"

	"github.com/F-Dupraz/Restauran-reservation-platform/reservations.git/models"
)

type ReservationRepository interface {
	CreateNewReservation(ctx context.Context, reservation *models.Reservation) error
}

var reservationImplementation ReservationRepository

func SetRepository(repository ReservationRepository) {
	reservationImplementation = repository
}

func CreateNewReservation(ctx context.Context, reservation *models.Reservation) error {
	return reservationImplementation.CreateNewReservation(ctx, reservation)
}
