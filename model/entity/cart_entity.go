package entity

import "time"

type Cart struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	UserID    int       `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	User      User      `gorm:"foreignKey:UserID"`
	CartItems []CartItem `gorm:"constraint:OnDelete:CASCADE"`
}
