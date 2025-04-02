package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                     uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Email                  string    `gorm:"type:varchar(255);not null;unique" json:"email"`
	Username               string    `gorm:"type:varchar(255)" json:"username"`
	Password               string    `gorm:"type:varchar(255)" json:"password"`
	IDRole                 string    `gorm:"type:varchar(50)" json:"id_role"`
	IDBusinessGroup        string    `gorm:"type:varchar(50)" json:"id_business_group"`
	IDSubBusinessGroup     string    `gorm:"type:varchar(50)" json:"id_sub_business_group"`
	EmailPIC               string    `gorm:"type:varchar(255)" json:"email_pic"`
	StatusActive           bool      `gorm:"type:bool" json:"status_active"`
	IDBusinessGroupDigital string    `gorm:"type:varchar(50)" json:"id_business_group_digital"`
	CreatedAt              time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt              time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// type MenuList struct {
// 	ID       uint   `gorm:"primaryKey"`
// 	CodeMenu string `gorm:"type:varchar(8);unique"`
// 	FormName string `gorm:"type:varchar(128)"`
// 	FormURL  string `gorm:"type:varchar(128)"`
// }

// type RoleMenu struct {
// 	ID     uint
// 	RoleID string `gorm:"type:varchar(6)"`
// 	MenuID uint
// }

// type UserMenu struct {
// 	ID        uint
// 	UserEmail string `gorm:"type:varchar(128)"`
// 	MenuID    uint
// }

// BeforeCreate untuk generate UUID sebelum menyimpan user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
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
}
