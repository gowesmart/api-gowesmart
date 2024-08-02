package entity

var RoleAdmin = "ADMIN"
var RoleUser = "USER"

type Profile struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	UserID uint   `gorm:"unique;not null"`
	Name   string `gorm:"type:varchar(100)"`
	Bio    string `gorm:"type:text"`
	Age    int    `gorm:"type:smallint"`
	Gender string `gorm:"type:varchar(6)" sql:"type:enum('MALE', 'FEMALE')"`
	User   User   `gorm:"foreignKey:UserID"`
}
