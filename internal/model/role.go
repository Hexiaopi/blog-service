package model

import "time"

type Role struct {
	ID         int       `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	Desc       string    `gorm:"column:desc" json:"desc"`
	State      uint8     `gorm:"state" json:"state"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

func (info Role) TableName() string {
	return "blog_role"
}
