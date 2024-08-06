package entity

import "time"

type CartItem struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	BikeID    uint      `gorm:"not null;uniqueIndex:idx_cart_bike"`
	CartID    uint      `gorm:"not null;uniqueIndex:idx_cart_bike"`
	Quantity  int       `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Cart      Cart      `gorm:"foreignKey:CartID"`
	Bike      Bike      `gorm:"foreignKey:BikeID"`
}