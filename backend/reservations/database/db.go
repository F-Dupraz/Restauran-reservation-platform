package database

import (
	"context"
	"database/sql"

	"github.com/F-Dupraz/Restauran-reservation-platform/reservations.git/models"
	_ "github.com/lib/pq"
)

type MyPostgresRepo struct {
	db *sql.DB
}

func NewMyPostgresRepo(url string) (*MyPostgresRepo, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &MyPostgresRepo{db}, nil
}

func (repo *MyPostgresRepo) Close() error {
	return repo.db.Close()
}

func (repo *MyPostgresRepo) CreateNewReservation(ctx context.Context, reservation *models.Reservation) error {
	_, err := repo.db.QueryContext(ctx, "INSERT INTO reservations (id, user_id, restaurant_id, day, num_guests) VALUES ($1, $2, $3, $4, $5);", reservation.Id, reservation.UserId, reservation.RestaurantId, reservation.Day, reservation.NumGuests)
	return err
}
