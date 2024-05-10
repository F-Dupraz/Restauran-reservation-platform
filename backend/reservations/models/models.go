package models

import "github.com/golang-jwt/jwt"

type Reservation struct {
	Id           string `json:"id"`
	UserId       string `json:"user_id"`
	RestaurantId string `json:"restaurant_id"`
	Day          []int  `json:"day"`
	From         string `json:"from"`
	To           string `json:"to"`
	NumGuests    int    `json:"num_guests"`
	IsDone       bool   `json:"is_done"`
}

type UserToken struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
