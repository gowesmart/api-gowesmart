package entity

import "time"

type Transaction struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	TotalPrice int       `gorm:"type:int;not null"`
	UserID     int       `gorm:"type:int;not null"`
	Status     string    `gorm:"type:varchar(255); not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	User       User      `gorm:"foreignKey:UserID"`
	Order      []Order   `gorm:"constraint:OnDelete:CASCADE"`
}
