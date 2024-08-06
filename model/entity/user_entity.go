package entity

import "time"

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	RoleID      uint   `gorm:"not null"`
	Username    string `gorm:"uniqueIndex:idx_username_email;not null;type:varchar(20)"`
	Email       string `gorm:"uniqueIndex:idx_username_email;not null;type:varchar(50)"`
	Password    string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Role        Role          `gorm:"foreignKey:RoleID"`
	Transaction []Transaction `gorm:"constraint:OnDelete:CASCADE"`
	Review      []Review      `gorm:"references:ID"`
}
