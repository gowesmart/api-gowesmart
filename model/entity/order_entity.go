package entity

import "time"

type Order struct {
	ID            int         `gorm:"primaryKey;autoIncrement"`
	BikeID        int         `gorm:"type:int;not null"`
	Quantity      int         `gorm:"type:int;not null"`
	TotalPrice    int         `gorm:"type:int;not null"`
	UserID        int         `gorm:"type:int;not null"`
	TransactionID int         `gorm:"type:int; not null"`
	CreatedAt     time.Time   `gorm:"autoCreateTime"`
	UpdatedAt     time.Time   `gorm:"autoUpdateTime"`
	User          User        `gorm:"foreignKey:UserID"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID"`
	Bike          Bike
}
