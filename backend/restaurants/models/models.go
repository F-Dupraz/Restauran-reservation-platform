package models

type Restaurant struct {
	Id string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	Owner string `json:"owner"`
	DaysOpen []string `json:"days_open:"`
	Specialties []string `json:"specialties"`
}
