package response

type OrderResponse struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	BikeID     int `gorm:"type:int;not null"`
	Quantity   int `gorm:"type:int;not null"`
	TotalPrice int `gorm:"type:int;not null"`
}

type GetAllOrderResponse struct {
	ID         int                     `json:"id"`
	Bike       GetAllOrderBikeResponse `json:"bike"`
	Quantity   int                     `json:"quantity"`
	TotalPrice int                     `json:"total_price"`
}

type GetAllOrderBikeResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}
