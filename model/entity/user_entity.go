package entity

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	RoleID    uint   `gorm:"not null"`
	Username  string `gorm:"unique;not null;type:varchar(20)"`
	Email     string `gorm:"unique;not null;type:varchar(50)"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Role      Role `gorm:"foreignKey:RoleID"`
}
