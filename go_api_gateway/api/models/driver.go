package models


type DriverResponse struct {
	Id string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	BirthDate string `json:"birth_date"`
}

type DriverRequest struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	BirthDate string `json:"birth_date"`
}

