package entity

import "time"

type Category struct {
	ID          uint   				`gorm:"primaryKey;autoIncrement"`
	Name    		string 				`gorm:"unique;not null;type:varchar(20)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Bike 				[]Bike 				`gorm:"constraint:OnDelete:CASCADE"`
}