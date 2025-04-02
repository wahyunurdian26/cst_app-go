package entity

type Role struct {
	ID          string `gorm:"type:varchar(6);primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(30)" json:"name"`
	Description string `gorm:"type:text" json:"desc"`
}
