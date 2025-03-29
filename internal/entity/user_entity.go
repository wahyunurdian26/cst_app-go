package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                     uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Email                  string         `gorm:"type:varchar(255);not null;unique" json:"email"`
	Username               string         `gorm:"type:varchar(255)" json:"username"`
	Password               string         `gorm:"type:varchar(255)" json:"password"`
	IDRole                 string         `gorm:"type:varchar(50)" json:"id_role"`
	IDBusinessGroup        uint           `gorm:"not null" json:"id_business_group"`
	IDSubBusinessGroup     uint           `gorm:"not null" json:"id_sub_business_group"`
	EmailPIC               string         `gorm:"type:varchar(255)" json:"email_pic"`
	StatusActive           string         `gorm:"type:varchar(50)" json:"status_active"`
	IDBusinessGroupDigital string         `gorm:"type:varchar(50)" json:"id_business_group_digital"`
	CreatedAt              time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt              time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// // Relasi ke SubGroup
	// SubGroup SubGroup `gorm:"foreignKey:IDSubBusinessGroup;references:ID" json:"sub_group"`

	// // Relasi ke Group melalui SubGroup
	// BusinessGroup Group `gorm:"foreignKey:IDBusinessGroup;references:ID" json:"business_group"`
}

type Group struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `gorm:"type:varchar(255);not null;unique" json:"name"`
	Prefix     string `gorm:"type:varchar(50)" json:"prefix"`
	QuotaMonth uint   `gorm:"not null" json:"quota_month"`
	QuotaDay   uint   `gorm:"not null" json:"quota_day"`
	CmpMonth   uint   `gorm:"not null" json:"cmp_month"`
	CmpDay     uint   `gorm:"not null" json:"cmp_day"`
}

type SubGroup struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name    string `gorm:"type:varchar(255);not null;unique" json:"name"`
	GroupID uint   `gorm:"not null" json:"group_id"`

	// 	// Relasi ke Group
	// 	Group Group `gorm:"foreignKey:GroupID;references:ID" json:"group"`
}
