package entity

type Menu struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	CodeMenu     string `gorm:"type:varchar(8);unique" json:"code_menu"`
	FormName     string `gorm:"type:varchar(128)" json:"form_name"`
	FormURL      string `gorm:"type:varchar(128)" json:"form_url"`
	FormCategory int    `gorm:"type:int" json:"form_category"`
	IDColor      int    `gorm:"type:int" json:"id_color"`
	IDIcon       int    `gorm:"type:int" json:"id_icon"`
}

type RoleMenu struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	RoleID string `gorm:"type:varchar(6)" json:"role_id"`
	MenuID uint   `gorm:"type:int" json:"menu_id"`
	Action int    `gorm:"type:int" json:"action"`
}

type UserMenu struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserEmail string `gorm:"type:varchar(128)" json:"user_email"`
	MenuID    uint   `gorm:"type:int" json:"menu_id"`
}
