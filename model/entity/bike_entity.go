package entity

import "time"

type Bike struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	CategoryID  uint   `gorm:"not null"`
	Name        string `gorm:"unique;not null;type:varchar(50)"`
	Brand       string `gorm:"not null;type:varchar(20)"`
	Description string `gorm:"type:varchar(1000)"`
	Year        int    `gorm:"not null"`
	Price       int    `gorm:"not null"`
	ImageUrl    string `gorm:"type:varchar(255)"`
	Stock       int    `gorm:"not null"`
	IsAvailable bool   `gorm:"not null;default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Category    Category `gorm:"foreignKey:CategoryID"`
	Review      []Review `gorm:"references:ID"`
}
