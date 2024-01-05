package model

import "time"

type RoleRest struct {
	ID         int       `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	RoleId     int       `gorm:"column:role_id" json:"role_id"`
	RestId     int       `gorm:"column:rest_id" json:"rest_id"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

func (rm *RoleRest) TableName() string {
	return "sys_role_rest"
}
