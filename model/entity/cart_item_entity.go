package entity

import "time"

type CartItem struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	BikeID    int       `gorm:"type:int;not null"`
	Quantity  int       `gorm:"type:int;not null"`
	Price     float64   `gorm:"type:float;not null"`
	CartID    int       `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Cart      Cart      `gorm:"foreignKey:CartID"`
	Bike      Bike      `gorm:"foreignKey:BikeID"`
}
