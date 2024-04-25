package database

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"github.com/F-Dupraz/Restauran-reservation-platform.git/models"
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

func (repo *MyPostgresRepo) InsterNewRestraurant(ctx context.Context, restaurant *models.Restaurant) error {
	DaysOpen := pq.Array(restaurant.DaysOpen)
	Specialties := pq.Array(restaurant.Specialties)
	_, err := repo.db.ExecContext(ctx, "INSERT INTO restaurants (id, name, city, owner, days_open, specialties) VALUES ($1, $2, $3, $4, $5, $6)", restaurant.Id, restaurant.Name, restaurant.City, restaurant.Owner, DaysOpen, Specialties)
	return err
}

func (repo *MyPostgresRepo) GetAllRestaurants(ctx context.Context, offset int) ([]models.Restaurant, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT name, city, days_open, specialties FROM restaurants ORDER BY created_at DESC LIMIT 20 OFFSET $1;", offset)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var restaurants = []models.Restaurant{}

	for rows.Next() {
		var restaurant = models.Restaurant{}
		var res_name string
		var res_city string
		var res_days sql.NullString
		var res_spec sql.NullString

		err = rows.Scan(&res_name, &res_city, &res_days, &res_spec)
		if err != nil {
			return nil, err
		}

		restaurant.Name = res_name
		restaurant.City = res_city

		var res_days_stringified = res_days.String
		res_days_stringified = strings.Replace(res_days.String, "{", "", -1)
		res_days_stringified = strings.Replace(res_days_stringified, "}", "", -1)
		restaurant.DaysOpen = strings.Split(res_days_stringified, ",")

		var res_spec_stringified = res_spec.String
		res_spec_stringified = strings.Replace(res_spec.String, "{", "", -1)
		res_spec_stringified = strings.Replace(res_spec_stringified, "}", "", -1)
		restaurant.Specialties = strings.Split(res_spec_stringified, ",")

		restaurants = append(restaurants, restaurant)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return restaurants, nil
}

func (repo *MyPostgresRepo) GetRestaurantByName(ctx context.Context, name string, offset int) ([]models.Restaurant, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT name, city, days_open, specialties FROM restaurants WHERE name = $1 ORDER BY created_at DESC LIMIT 20 OFFSET $2;", strings.ToLower(name), offset)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var restaurants = []models.Restaurant{}

	for rows.Next() {
		var restaurant = models.Restaurant{}
		var res_name string
		var res_city string
		var res_days sql.NullString
		var res_spec sql.NullString

		err = rows.Scan(&res_name, &res_city, &res_days, &res_spec)
		if err != nil {
			return nil, err
		}

		restaurant.Name = res_name
		restaurant.City = res_city

		var res_days_stringified = res_days.String
		res_days_stringified = strings.Replace(res_days.String, "{", "", -1)
		res_days_stringified = strings.Replace(res_days_stringified, "}", "", -1)
		restaurant.DaysOpen = strings.Split(res_days_stringified, ",")

		var res_spec_stringified = res_spec.String
		res_spec_stringified = strings.Replace(res_spec.String, "{", "", -1)
		res_spec_stringified = strings.Replace(res_spec_stringified, "}", "", -1)
		restaurant.Specialties = strings.Split(res_spec_stringified, ",")

		restaurants = append(restaurants, restaurant)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return restaurants, nil
}

func (repo *MyPostgresRepo) GetRestaurantByCity(ctx context.Context, city string, offset int) ([]models.Restaurant, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT name, city, days_open, specialties FROM restaurants WHERE city = $1 ORDER BY created_at DESC LIMIT 20 OFFSET $2;", strings.ToLower(city), offset)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var restaurants = []models.Restaurant{}

	for rows.Next() {
		var restaurant = models.Restaurant{}
		var res_name string
		var res_city string
		var res_days sql.NullString
		var res_spec sql.NullString

		err = rows.Scan(&res_name, &res_city, &res_days, &res_spec)
		if err != nil {
			return nil, err
		}

		restaurant.Name = res_name
		restaurant.City = res_city

		var res_days_stringified = res_days.String
		res_days_stringified = strings.Replace(res_days.String, "{", "", -1)
		res_days_stringified = strings.Replace(res_days_stringified, "}", "", -1)
		restaurant.DaysOpen = strings.Split(res_days_stringified, ",")

		var res_spec_stringified = res_spec.String
		res_spec_stringified = strings.Replace(res_spec.String, "{", "", -1)
		res_spec_stringified = strings.Replace(res_spec_stringified, "}", "", -1)
		restaurant.Specialties = strings.Split(res_spec_stringified, ",")

		restaurants = append(restaurants, restaurant)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return restaurants, nil
}

// -------------------- TERMINAR -------------------- //

// func (repo *MyPostgresRepo) UpdateRestaurant(ctx context.Context, restaurant_data *models.Restaurant,  id int64) error {
// 	_, err := repo.db.ExecContext(ctx, "")
// 	return err
// }

// func (repo *MyPostgresRepo) DeleteRestaurant(ctx context.Context, id string) error {
// 	_, err := repo.db.ExecContext(ctx, "DELETE FROM restaurants WHERE id = $1", id)
// 	return err
// }
