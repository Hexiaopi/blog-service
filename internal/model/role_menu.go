package model

import "time"

type RoleMenu struct {
	ID         int       `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	RoleId     int       `gorm:"column:role_id" json:"role_id"`
	MenuId     int       `gorm:"column:menu_id" json:"menu_id"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

func (rm *RoleMenu) TableName() string {
	return "sys_role_menu"
}
