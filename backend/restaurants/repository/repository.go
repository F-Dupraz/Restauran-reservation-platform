package repository

import (
	"context"

	"github.com/F-Dupraz/Restauran-reservation-platform.git/models"
)

type RestaurantRepository interface {
	InsterNewRestraurant(ctx context.Context, restaurant *models.Restaurant) error
	GetAllRestaurants(ctx context.Context, offset int) ([]models.Restaurant, error)
	GetRestaurantByName(ctx context.Context, name string, offset int) ([]models.Restaurant, error)
	GetRestaurantByCity(ctx context.Context, city string, offset int) ([]models.Restaurant, error)
	// UpdateRestaurant(ctx context.Context, id string) error
	// DeleteRestaurant(ctx context.Context, id string) error
	Close() error
}

var restaurantImplementation RestaurantRepository

func SetRepository(repository RestaurantRepository) {
	restaurantImplementation = repository
}

func InsterNewRestraurant(ctx context.Context, restaurant *models.Restaurant) error {
	return restaurantImplementation.InsterNewRestraurant(ctx, restaurant)
}

func GetAllRestaurants(ctx context.Context, offset int) ([]models.Restaurant, error) {
	return restaurantImplementation.GetAllRestaurants(ctx, offset)
}

func GetRestaurantByName(ctx context.Context, name string, offset int) ([]models.Restaurant, error) {
	return restaurantImplementation.GetRestaurantByName(ctx, name, offset)
}

func GetRestaurantByCity(ctx context.Context, city string, offset int) ([]models.Restaurant, error) {
	return restaurantImplementation.GetRestaurantByCity(ctx, city, offset)
}

// func UpdateRestaurant(ctx context.Context, id string) error {
// 	return restaurantImplementation.UpdateRestaurant(ctx, id)
// }

// func DeleteRestaurant(ctx context.Context, id string) error {
// 	return restaurantImplementation.DeleteRestaurant(ctx, id)
// }

func Close() error {
	return restaurantImplementation.Close()
}
