package model

import "time"

type SysRest struct {
	ID         int       `gorm:"id" json:"id"`
	Name       string    `gorm:"name" json:"name"`
	Path       string    `gorm:"path" json:"path"`
	Method     string    `gorm:"method" json:"method"`
	ParentId   int       `gorm:"parent_id" json:"parent_id"`
	CreateTime time.Time `gorm:"create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"update_time" json:"update_time"`
}

func (i *SysRest) TableName() string {
	return "sys_rest"
}
