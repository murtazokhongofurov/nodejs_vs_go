package models


type CarRequest struct {
	DriverId string `json:"driver_id"`
	CarModel string `json:"car_model"`
	ManufactureYear string `json:"manufacture_year"`
	NumberPlate string `json:"number_plate"`
	TechnicalLicence string `json:"technical_licence"`
	CarType string `json:"car_type"`
	Color string `json:"color"`
	CarImage string `json:"car_image"`
}


type CarResponse struct {
	Id string `json:"id"`
	DriverId string `json:"driver_id"`
	CarModel string `json:"car_model"`
	ManufactureYear string `json:"manufacture_year"`
	NumberPlate string `json:"number_plate"`
	TechnicalLicence string `json:"technical_licence"`
	CarType string `json:"car_type"`
	Color string `json:"color"`
	CarImage string `json:"car_image"`
}
