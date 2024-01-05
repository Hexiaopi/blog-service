package model

import "time"

type SysMenu struct {
	ID         int       `gorm:"id" json:"id"`
	Name       string    `gorm:"name" json:"name"`
	Path       string    `gorm:"path" json:"path"`
	Title      string    `gorm:"title" json:"title"`
	Icon       string    `gorm:"icon" json:"icon"`
	Component  string    `gorm:"component" json:"component"`
	Sort       int       `gorm:"order" json:"sort"`
	Redirect   string    `gorm:"redirect" json:"redirect"`
	Hidden     bool      `gorm:"hidden" json:"is_hidden"`
	ParentId   int       `gorm:"parent_id" json:"parent_id"`
	CreateTime time.Time `gorm:"create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"update_time" json:"update_time"`
}

func (i *SysMenu) TableName() string {
	return "sys_menu"
}
