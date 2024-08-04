package entity

import "time"

type Review struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Comment   string    `gorm:"type:text;not null" json:"comment"`
	Rating    int       `gorm:"not null" json:"rating"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	BikeID    uint      `gorm:"not null" json:"bike_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
}
