package entity

type Group struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `gorm:"type:varchar(255);not null;unique" json:"name"`
	Prefix     string `gorm:"type:varchar(50)" json:"prefix"`
	QuotaMonth uint   `gorm:"not null" json:"quota_month"`
	QuotaDay   uint   `gorm:"not null" json:"quota_day"`
	CmpMonth   uint   `gorm:"not null" json:"cmp_month"`
	CmpDay     uint   `gorm:"not null" json:"cmp_day"`
}
