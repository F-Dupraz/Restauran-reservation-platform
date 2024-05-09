package models

import "github.com/golang-jwt/jwt"

type Restaurant struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Owner        string   `json:"owner"`
	Address      string   `json:"address"`
	Description  string   `json:"description"`
	City         string   `json:"city"`
	DaysOpen     []int    `json:"days_open"`
	WorkingHours []string `json:"working_hours"`
	Capacity     []int    `json:"capacity"`
	Specialties  []string `json:"specialties"`
}

type UserToken struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
