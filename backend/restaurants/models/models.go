package models

import "github.com/golang-jwt/jwt"

type Restaurant struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	City        string   `json:"city"`
	Owner       string   `json:"owner"`
	Address     string   `json:"address"`
	Description string   `json:"description"`
	DaysOpen    []string `json:"days_open"`
	Capacity    []int    `json:"capacity"`
	Specialties []string `json:"specialties"`
}

type UserToken struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
