package response

import "time"

type BikeResponse struct {
	ID          uint      `json:"id"`
	CategoryID  uint      `json:"category_id"`
	Name        string    `json:"name"`
	Brand       string    `json:"brand"`
	Description string    `json:"description"`
	Year        int       `json:"year"`
	Price       int       `json:"price"`
	ImageUrl    string    `json:"image_url"`
	Stock       int       `json:"stock"`
	IsAvailable bool      `json:"is_available"`
	Rating      int       `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BikeListResponse struct {
	Bikes []BikeResponse `json:"bikes"`
}
