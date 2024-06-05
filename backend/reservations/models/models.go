package models

import (
	"github.com/golang-jwt/jwt"
)

type Reservation struct {
	Id           string `json:"id"`
	UserId       string `json:"user_id"`
	RestaurantId string `json:"restaurant_id"`
	Day          string `json:"day"`
	DayInt       []int  `json:"day_int"`
	From         string `json:"from"`
	To           string `json:"to"`
	NumGuests    int    `json:"num_guests"`
	IsDone       bool   `json:"is_done"`
}

type AReservation struct {
	Id             string `json:"id"`
	RestaurantName string `json:"restaurant_name"`
	Day            string `json:"day"`
	From           string `json:"from"`
	To             string `json:"to"`
	NumGuests      int    `json:"num_guests"`
}

type MyReservation struct {
	Id      string `json:"id"`
	Day     string `json:"day"`
	ResName string `json:"restaurant_name"`
}

type UserToken struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
