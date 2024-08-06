package response

type CartItemResponse struct {
	ID       int     `json:"id"`
	CartID   int     `json:"cart_id"`
	BikeID   int     `json:"bike_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
