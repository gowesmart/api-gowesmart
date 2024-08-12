package response

import "time"

type CartItemResponse struct {
	ID        uint      `json:"id"`
	CartID    uint      `json:"cart_id"`
	BikeID    uint      `json:"bike_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type GetUserCartItemResponse struct {
	ID        uint                        `json:"id"`
	CartID    uint                        `json:"cart_id"`
	Bike      GetUserCartItemBikeResponse `json:"bike"`
	Quantity  int                         `json:"quantity"`
	Price     float64                     `json:"price,omitempty"`
	CreatedAt time.Time                   `json:"created_at,omitempty"`
	UpdatedAt time.Time                   `json:"updated_at,omitempty"`
}

type GetUserCartItemBikeResponse struct {
	ID          uint   `json:"id"`
	Brand       string `json:"brand"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	ImageUrl    string `json:"image_url"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
}
