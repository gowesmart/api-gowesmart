package entity

import "time"

type CartItem struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	BikeID    uint      `gorm:"not null"`
	CartID    uint      `gorm:"not null"`
	Quantity  int       `gorm:"type:int;not null"`
	Price     float64   `gorm:"type:float;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Cart      Cart      `gorm:"foreignKey:CartID"`
	Bike      Bike      `gorm:"foreignKey:BikeID"`
}
