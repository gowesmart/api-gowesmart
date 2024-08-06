package request

type CartItemCreateRequest struct {
	BikeID   uint `json:"bike_id" binding:"required"`
	Quantity int  `json:"quantity" binding:"required"`
}

type CartItemUpdateRequest struct {
	BikeID   uint `json:"bike_id" binding:"required"`
	Quantity int  `json:"quantity" binding:"required"`
}

type CartItemDeleteRequest struct {
	BikeID uint `json:"bike_id" binding:"required"`
}
