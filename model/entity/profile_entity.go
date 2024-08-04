package entity

type Profile struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	UserID uint   `gorm:"unique;not null"`
	Name   string `gorm:"type:varchar(100);default:null"`
	Bio    string `gorm:"type:text;default:null"`
	Age    int    `gorm:"type:smallint;default:null"`
	Gender string `gorm:"type:varchar(6);default:null" sql:"type:enum('MALE', 'FEMALE')"`
	User   User   `gorm:"foreignKey:UserID"`
}
