package entity

import "time"

type Cart struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	UserID    uint       `gorm:"unique;not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	User      User       `gorm:"foreignKey:UserID"`
	CartItem []CartItem 
}