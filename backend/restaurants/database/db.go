package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/F-Dupraz/Restauran-reservation-platform.git/models"
)

type MySQLRepo struct {
	db *sql.DB
}

func NewMySQLRepo(url string) (*MySQLRepo, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &MySQLRepo{db}, nil
}

func (repo *MySQLRepo) Close() error {
	return repo.db.Close()
}

func (repo *MySQLRepo) InsterNewRestraurant(ctx context.Context, restaurant *models.Restaurant) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO restaurants (name, city, owner, days_open, specialties) VALUES ($1, $2, $3, $4, $5)", restaurant.Name, restaurant.City, restaurant.Owner, restaurant.days_open, restaurant.Specialties)
	return err
}

func (repo *MySQLRepo) GetRestaurantByName(ctx context.Context, name string) ([]*models.Restaurant, error) {
	rows, err := repo.db.ExecContext(ctx, "SELECT name, city, days_open, specialties FROM restaurants WHERE name = $1", name)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var restaurants = []models.Restaurant
	for rows.Next() {
		if err = rows.Scan(&restaurants.Name, &restaurants.City, &restaurants.DaysOpen, &restaurants.Specialties); err == nil {
			return &restaurants, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &restaurants, nil
}

func (repo *MySQLRepo) GetRestaurantNyCity(ctx context.Context, city string) ([]*models.Restaurant, error) {
	rows, err := repo.db.ExecContext(ctx, "SELECT name, city, days_open, specialties FROM restaurants WHERE city = $1", city)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var restaurants = []models.Restaurant
	for rows.Next() {
		if err = rows.Scan(&restaurants.Name, &restaurants.City, &restaurants.DaysOpen, &restaurants.Specialties); err == nil {
			return &restaurants, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &restaurants, nil
}

// -------------------- TERMINAR -------------------- //

// func (repo *MySQLRepo) UpdateRestaurant(ctx context.Context, restaurant_data *models.Restaurant,  id int64) error {
// 	_, err := repo.db.ExecContext(ctx, "")
// 	return err
// }

func (repo *MySQLRepo) DeleteRestaurant(ctx context.Context, id string) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM restaurants WHERE id = $1", id)
	return err
}
