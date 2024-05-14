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
	WorkingHours := pq.Array(restaurant.WorkingHours)
	Capacity := pq.Array(restaurant.Capacity)
	Specialties := pq.Array(restaurant.Specialties)
	_, err := repo.db.QueryContext(ctx, "INSERT INTO restaurants (id, name, city, owner, address, description, days_open, working_hours, capacity, specialties) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);", restaurant.Id, restaurant.Name, restaurant.City, restaurant.Owner, restaurant.Address, restaurant.Description, DaysOpen, WorkingHours, Capacity, Specialties)
	return err
}

func (repo *MyPostgresRepo) GetAllRestaurants(ctx context.Context, offset int) ([]models.Restaurant, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, city, owner, address, description, days_open, working_hours, capacity, specialties FROM restaurants ORDER BY created_at DESC LIMIT 20 OFFSET $1;", offset)
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
		var res_id string
		var res_name string
		var res_city string
		var res_owner string
		var res_address string
		var res_des string
		var res_days pq.Int64Array
		var res_wh sql.NullString
		var res_cap pq.Int64Array
		var res_spec sql.NullString

		err = rows.Scan(&res_id, &res_name, &res_city, &res_owner, &res_address, &res_des, &res_days, &res_wh, &res_cap, &res_spec)
		if err != nil {
			return nil, err
		}

		restaurant.Id = res_id
		restaurant.Name = res_name
		restaurant.City = res_city
		restaurant.Owner = res_owner
		restaurant.Address = res_address
		restaurant.Description = res_des

		restaurant.Capacity = make([]int, len(res_cap))
		for i, v := range res_cap {
			restaurant.Capacity[i] = int(v)
		}

		var res_wh_stringified = res_wh.String
		res_wh_stringified = strings.Replace(res_wh.String, "{", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "}", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "[", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "]", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "\\", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "\"", "", -1)
		restaurant.WorkingHours = strings.Split(res_wh_stringified, ",")

		restaurant.DaysOpen = make([]int, len(res_days))
		for i, v := range res_days {
			restaurant.DaysOpen[i] = int(v)
		}

		var res_spec_stringified = res_spec.String
		res_spec_stringified = strings.Replace(res_spec.String, "{", "", -1)
		res_spec_stringified = strings.Replace(res_spec_stringified, "}", "", -1)
		res_spec_stringified = strings.Replace(res_spec_stringified, "\\", "", -1)
		res_spec_stringified = strings.Replace(res_spec_stringified, "\"", "", -1)
		restaurant.Specialties = strings.Split(res_spec_stringified, ",")

		restaurants = append(restaurants, restaurant)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return restaurants, nil
}

func (repo *MyPostgresRepo) GetRestaurantByName(ctx context.Context, name string, offset int) ([]models.Restaurant, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, city, owner, address, description, days_open, working_hours, capacity, specialties FROM restaurants WHERE name = $1 ORDER BY created_at DESC LIMIT 20 OFFSET $2;", strings.ToLower(name), offset)
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
		var res_id string
		var res_name string
		var res_city string
		var res_owner string
		var res_address string
		var res_des string
		var res_days pq.Int64Array
		var res_wh sql.NullString
		var res_cap pq.Int64Array
		var res_spec sql.NullString

		err = rows.Scan(&res_id, &res_name, &res_city, &res_owner, &res_address, &res_des, &res_days, &res_wh, &res_cap, &res_spec)
		if err != nil {
			return nil, err
		}

		restaurant.Id = res_id
		restaurant.Name = res_name
		restaurant.City = res_city
		restaurant.Owner = res_owner
		restaurant.Address = res_address
		restaurant.Description = res_des

		restaurant.Capacity = make([]int, len(res_cap))
		for i, v := range res_cap {
			restaurant.Capacity[i] = int(v)
		}

		var res_wh_stringified = res_wh.String
		res_wh_stringified = strings.Replace(res_wh.String, "{", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "}", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "[", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "]", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "\\", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "\"", "", -1)
		restaurant.WorkingHours = strings.Split(res_wh_stringified, ",")

		restaurant.DaysOpen = make([]int, len(res_days))
		for i, v := range res_days {
			restaurant.DaysOpen[i] = int(v)
		}

		var res_spec_stringified = res_spec.String
		res_spec_stringified = strings.Replace(res_spec.String, "{", "", -1)
		res_spec_stringified = strings.Replace(res_spec_stringified, "}", "", -1)
		res_spec_stringified = strings.Replace(res_spec_stringified, "\\", "", -1)
		res_spec_stringified = strings.Replace(res_spec_stringified, "\"", "", -1)
		restaurant.Specialties = strings.Split(res_spec_stringified, ",")

		restaurants = append(restaurants, restaurant)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return restaurants, nil
}

func (repo *MyPostgresRepo) GetRestaurantByCity(ctx context.Context, city string, offset int) ([]models.Restaurant, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, city, owner, address, description, days_open, working_hours, capacity, specialties FROM restaurants WHERE city = $1 ORDER BY created_at DESC LIMIT 20 OFFSET $2;", strings.ToLower(city), offset)
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
		var res_id string
		var res_name string
		var res_city string
		var res_owner string
		var res_address string
		var res_des string
		var res_days pq.Int64Array
		var res_wh sql.NullString
		var res_cap pq.Int64Array
		var res_spec sql.NullString

		err = rows.Scan(&res_id, &res_name, &res_city, &res_owner, &res_address, &res_des, &res_days, &res_wh, &res_cap, &res_spec)
		if err != nil {
			return nil, err
		}

		restaurant.Id = res_id
		restaurant.Name = res_name
		restaurant.City = res_city
		restaurant.Owner = res_owner
		restaurant.Address = res_address
		restaurant.Description = res_des

		restaurant.Capacity = make([]int, len(res_cap))
		for i, v := range res_cap {
			restaurant.Capacity[i] = int(v)
		}

		var res_wh_stringified = res_wh.String
		res_wh_stringified = strings.Replace(res_wh.String, "{", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "}", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "[", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "]", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "\\", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "\"", "", -1)
		restaurant.WorkingHours = strings.Split(res_wh_stringified, ",")

		restaurant.DaysOpen = make([]int, len(res_days))
		for i, v := range res_days {
			restaurant.DaysOpen[i] = int(v)
		}

		var res_spec_stringified = res_spec.String
		res_spec_stringified = strings.Replace(res_spec.String, "{", "", -1)
		res_spec_stringified = strings.Replace(res_spec_stringified, "}", "", -1)
		res_spec_stringified = strings.Replace(res_spec_stringified, "\\", "", -1)
		res_spec_stringified = strings.Replace(res_spec_stringified, "\"", "", -1)
		restaurant.Specialties = strings.Split(res_spec_stringified, ",")

		restaurants = append(restaurants, restaurant)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return restaurants, nil
}

func (repo *MyPostgresRepo) GetMyRestaurants(ctx context.Context, id string) ([]models.Restaurant, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, city, owner, address, description, days_open, working_hours, capacity, specialties FROM restaurants WHERE owner = $1 ORDER BY created_at DESC;", id)
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
		var res_id string
		var res_name string
		var res_city string
		var res_owner string
		var res_address string
		var res_des string
		var res_days pq.Int64Array
		var res_wh sql.NullString
		var res_cap pq.Int64Array
		var res_spec sql.NullString

		err = rows.Scan(&res_id, &res_name, &res_city, &res_owner, &res_address, &res_des, &res_days, &res_wh, &res_cap, &res_spec)
		if err != nil {
			return nil, err
		}

		restaurant.Id = res_id
		restaurant.Name = res_name
		restaurant.City = res_city
		restaurant.Owner = res_owner
		restaurant.Address = res_address
		restaurant.Description = res_des

		restaurant.Capacity = make([]int, len(res_cap))
		for i, v := range res_cap {
			restaurant.Capacity[i] = int(v)
		}

		var res_wh_stringified = res_wh.String
		res_wh_stringified = strings.Replace(res_wh.String, "{", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "}", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "[", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "]", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "\\", "", -1)
		res_wh_stringified = strings.Replace(res_wh_stringified, "\"", "", -1)
		restaurant.WorkingHours = strings.Split(res_wh_stringified, ",")

		restaurant.DaysOpen = make([]int, len(res_days))
		for i, v := range res_days {
			restaurant.DaysOpen[i] = int(v)
		}

		var res_spec_stringified = res_spec.String
		res_spec_stringified = strings.Replace(res_spec.String, "{", "", -1)
		res_spec_stringified = strings.Replace(res_spec_stringified, "}", "", -1)
		res_spec_stringified = strings.Replace(res_spec_stringified, "\\", "", -1)
		res_spec_stringified = strings.Replace(res_spec_stringified, "\"", "", -1)
		restaurant.Specialties = strings.Split(res_spec_stringified, ",")

		restaurants = append(restaurants, restaurant)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return restaurants, nil
}

func (repo *MyPostgresRepo) UpdateRestaurant(ctx context.Context, restaurant_data *models.Restaurant, user_id string) error {
	DaysOpen := pq.Array(restaurant_data.DaysOpen)
	WorkingHours := pq.Array(restaurant_data.WorkingHours)
	Capacity := pq.Array(restaurant_data.Capacity)
	Specialties := pq.Array(restaurant_data.Specialties)
	_, err := repo.db.ExecContext(ctx, "UPDATE restaurants SET name=$1, description=$2, days_open=$3, working_hours=$4, capacity=$5, specialties=$6 WHERE id=$7 AND owner=$8;", restaurant_data.Name, restaurant_data.Address, DaysOpen, WorkingHours, Capacity, Specialties, restaurant_data.Id, user_id)
	return err
}

func (repo *MyPostgresRepo) DeleteRestaurant(ctx context.Context, id string, user_id string) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM restaurants WHERE id = $1 AND owner = $2", id, user_id)
	return err
}
