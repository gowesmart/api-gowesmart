package request

type CartItemCreateRequest struct {
	CartID   uint    `json:"cart_id" binding:"required"`
	BikeID   uint    `json:"bike_id" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
}

type CartItemUpdateRequest struct {
	ID       uint    `json:"id" binding:"required"`
	CartID   uint    `json:"cart_id" binding:"required"`
	BikeID   uint    `json:"bike_id" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
}
