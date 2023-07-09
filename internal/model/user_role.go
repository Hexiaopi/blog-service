package model

import "time"

type UserRole struct {
	ID         int       `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	UserId     int       `gorm:"column:user_id" json:"user_id"`
	RoleId     int       `gorm:"column:role_id" json:"role_id"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

func (ur *UserRole) TableName() string {
	return "sys_user_role"
}
