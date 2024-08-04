package response

type OrderResponse struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	BikeID     int `gorm:"type:int;not null"`
	Quantity   int `gorm:"type:int;not null"`
	TotalPrice int `gorm:"type:int;not null"`
}
