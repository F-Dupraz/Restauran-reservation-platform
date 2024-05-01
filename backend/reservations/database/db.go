package database

import (
	"context"
	"database/sql"
	"log"

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

func (repo *MyPostgresRepo) UpdateReservation(ctx context.Context, reservation *models.Reservation, id string) error {
	_, err := repo.db.QueryContext(ctx, "UPDATE reservations SET day=$1, num_guests=$2 WHERE id=$3 AND user_id=$4;", reservation.Day, reservation.NumGuests, reservation.Id, id)
	return err
}

func (repo *MyPostgresRepo) GetReservationById(ctx context.Context, id string, user_id string) (*models.Reservation, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, user_id, restaurant_id, day, num_guests, is_done FROM reservations WHERE id=$1 AND user_id=$2;", id, user_id)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var reservation = models.Reservation{}

	for rows.Next() {
		err = rows.Scan(&reservation.Id, &reservation.UserId, &reservation.RestaurantId, &reservation.Day, &reservation.NumGuests, &reservation.IsDone)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &reservation, nil
}
