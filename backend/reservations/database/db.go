package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/F-Dupraz/Restauran-reservation-platform/reservations.git/models"
	"github.com/lib/pq"
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
	Day := pq.Array(reservation.Day)
	_, err := repo.db.QueryContext(ctx, "INSERT INTO reservations (id, user_id, restaurant_id, day, h_from, h_to, num_guests) VALUES ($1, $2, $3, $4, $5, $6, $7);", reservation.Id, reservation.UserId, reservation.RestaurantId, Day, reservation.From, reservation.To, reservation.NumGuests)
	return err
}

func (repo *MyPostgresRepo) UpdateReservation(ctx context.Context, reservation *models.Reservation, id string) error {
	Day := pq.Array(reservation.Day)
	_, err := repo.db.QueryContext(ctx, "UPDATE reservations SET day=$1, h_from=$2, h_to=$3, num_guests=$4 WHERE id=$5 AND user_id=$6;", Day, reservation.From, reservation.To, reservation.NumGuests, reservation.Id, id)
	return err
}

func (repo *MyPostgresRepo) GetReservationById(ctx context.Context, id string) (*models.Reservation, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, user_id, restaurant_id, day, h_from, h_to, num_guests, is_done FROM reservations WHERE id=$1;", id)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var reservation_day pq.Int64Array

	var reservation = models.Reservation{}

	for rows.Next() {
		err = rows.Scan(&reservation.Id, &reservation.UserId, &reservation.RestaurantId, &reservation_day, &reservation.From, &reservation.To, &reservation.NumGuests, &reservation.IsDone)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	reservation.Day = make([]int, len(reservation_day))
	for i, v := range reservation_day {
		reservation.Day[i] = int(v)
	}

	return &reservation, nil
}

func (repo *MyPostgresRepo) GetReservationsByDay(ctx context.Context, day [1]int, restaurant_id string) ([]models.Reservation, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, user_id, restaurant_id, day, h_from, h_to, num_guests, is_done FROM reservations WHERE day[1]=$1 AND restaurant_id=$2;", day[0], restaurant_id)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var reservations = []models.Reservation{}

	for rows.Next() {
		var reservation = models.Reservation{}

		var res_id string
		var res_uid string
		var res_rid string
		var res_from string
		var res_to string
		var res_day pq.Int64Array
		var res_nguest int
		var res_isdone bool

		err = rows.Scan(&res_id, &res_uid, &res_rid, &res_day, &res_from, &res_to, &res_nguest, &res_isdone)

		reservation.Day = make([]int, len(res_day))
		for i, v := range res_day {
			reservation.Day[i] = int(v)
		}

		reservation.Id = res_id
		reservation.UserId = res_uid
		reservation.RestaurantId = res_rid
		reservation.NumGuests = res_nguest
		reservation.IsDone = res_isdone

		reservations = append(reservations, reservation)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}

func (repo *MyPostgresRepo) GetReservationsByRid(ctx context.Context, restaurant_id string) ([]models.Reservation, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, user_id, restaurant_id, day, h_from, h_to, num_guests, is_done FROM reservations WHERE restaurant_id=$1;", restaurant_id)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var reservations = []models.Reservation{}

	for rows.Next() {
		var reservation = models.Reservation{}

		var res_id string
		var res_uid string
		var res_rid string
		var res_from string
		var res_to string
		var res_day pq.Int64Array
		var res_nguest int
		var res_isdone bool

		err = rows.Scan(&res_id, &res_uid, &res_rid, &res_day, &res_from, &res_to, &res_nguest, &res_isdone)

		reservation.Day = make([]int, len(res_day))
		for i, v := range res_day {
			reservation.Day[i] = int(v)
		}

		reservation.Id = res_id
		reservation.UserId = res_uid
		reservation.RestaurantId = res_rid
		reservation.NumGuests = res_nguest
		reservation.IsDone = res_isdone

		reservations = append(reservations, reservation)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}
