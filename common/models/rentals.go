package models

type Rental struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Type            string  `json:"type"`
	VehicleMake     string  `json:"make"`
	VehicleModel    string  `json:"model"`
	VehicleYear     int     `json:"year"`
	VehicleLength   float64 `json:"length"`
	Sleeps          int     `json:"sleeps"`
	PrimaryImageUrl string  `json:"primary_image_url"`
	PricePerDay     int     `json:"price_per_day"`
	HomeCity        string  `json:"city"`
	HomeState       string  `json:"state"`
	HomeCountry     string  `json:"country"`
	HomeZip         string  `json:"zip"`
	Lat             float64 `json:"lat"`
	Lng             float64 `json:"lng"`
	UserId          int     `json:"user_id"`
	User            User    `json:"user"`
}
