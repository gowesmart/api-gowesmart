package entity

import "time"

type Cart struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	UserID    uint       `gorm:"not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	User      User       `gorm:"foreignKey:UserID"`
	CartItems []CartItem `gorm:"constraint:OnDelete:CASCADE"`
}
