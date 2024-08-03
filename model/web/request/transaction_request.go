package request

type TransactionCreate struct {
	BikeID     int `json:"bikeId" bind:"required"`
	Quantity   int `json:"quantity" bind:"required"`
	TotalPrice int `json:"totalPrice" bind:"required"`
}

type TransactionUpdate struct {
	ID         int `json:"id" bind:"required"`
	BikeID     int `json:"bikeId" bind:"required"`
	Quantity   int `json:"quantity" bind:"required"`
	TotalPrice int `json:"totalPrice" bind:"required"`
}