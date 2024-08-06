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
