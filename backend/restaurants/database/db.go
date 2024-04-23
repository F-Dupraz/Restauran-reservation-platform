package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/F-Dupraz/Restauran-reservation-platform.git/models"
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

func (repo *MyPostgresRepo) InsterNewRestraurant(ctx context.Context, restaurant *models.Restaurant) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO restaurants (name, city, owner, days_open, specialties) VALUES ($1, $2, $3, $4, $5)", restaurant.Name, restaurant.City, restaurant.Owner, restaurant.DaysOpen, restaurant.Specialties)
	return err
}

func (repo *MyPostgresRepo) GetRestaurantByName(ctx context.Context, name string) ([]models.Restaurant, error) {
	_, err := repo.db.ExecContext(ctx, "SELECT name, city, days_open, specialties FROM restaurants WHERE name = $1", name)
	if err != nil {
		return nil, err
	}

	var restaurants = []models.Restaurant{}

	fmt.Println(restaurants)

	return restaurants, nil

	// for rows.Next() {
	// 	if err = rows.Scan(&restaurants.Name, &restaurants.City, &restaurants.DaysOpen, &restaurants.Specialties); err == nil {
	// 		return &restaurants, nil
	// 	}
	// }
	// if err = rows.Err(); err != nil {
	// 	return nil, err
	// }
	// return &restaurants, nil
}

func (repo *MyPostgresRepo) GetRestaurantByCity(ctx context.Context, city string) ([]models.Restaurant, error) {
	_, err := repo.db.ExecContext(ctx, "SELECT name, city, days_open, specialties FROM restaurants WHERE city = $1", city)
	if err != nil {
		return nil, err
	}

	var restaurants = []models.Restaurant{}

	fmt.Println(restaurants)

	return restaurants, nil

	// for rows.Next() {
	// 	if err = rows.Scan(&restaurants.Name, &restaurants.City, &restaurants.DaysOpen, &restaurants.Specialties); err == nil {
	// 		return &restaurants, nil
	// 	}
	// }
	// if err = rows.Err(); err != nil {
	// 	return nil, err
	// }
	// return &restaurants, nil
}

// -------------------- TERMINAR -------------------- //

// func (repo *MyPostgresRepo) UpdateRestaurant(ctx context.Context, restaurant_data *models.Restaurant,  id int64) error {
// 	_, err := repo.db.ExecContext(ctx, "")
// 	return err
// }

func (repo *MyPostgresRepo) DeleteRestaurant(ctx context.Context, id string) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM restaurants WHERE id = $1", id)
	return err
}
