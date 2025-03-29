package entity

type SubGroup struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name    string `gorm:"type:varchar(255);not null;unique" json:"name"`
	GroupID uint   `gorm:"not null" json:"group_id"`
}
