package response

type CategoryResponse struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"unique;not null;type:varchar(20)"`
}