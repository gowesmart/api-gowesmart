package request

type TransactionCreate struct {
	BikeID     int `json:"bike_id" bind:"required"`
	Quantity   int `json:"quantity" bind:"required"`
	TotalPrice int `json:"total_price" bind:"required"`
}

type TransactionUpdate struct {
	ID         int `json:"id" bind:"required"`
	BikeID     int `json:"bike_id" bind:"required"`
	Quantity   int `json:"quantity" bind:"required"`
	TotalPrice int `json:"total_price" bind:"required"`
}
