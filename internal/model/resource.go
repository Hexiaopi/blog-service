package model

import "time"

type Resource struct {
	ID         int       `gorm:"primaryKey;autoIncrement;id" json:"id"`
	Name       string    `gorm:"name" json:"name"`
	Blob       string    `gorm:"blob" json:"blob"`
	Type       string    `gorm:"type" json:"type"`
	Size       int64     `gorm:"size" json:"size"`
	State      uint8     `gorm:"state" json:"state"`
	CreateTime time.Time `gorm:"create_time" json:"create_time,omitempty"`
	UpdateTime time.Time `gorm:"update_time" json:"update_time,omitempty"`
	Operator   string    `gorm:"operator" json:"operator"`
}

func (r Resource) TableName() string {
	return "blog_resource"
}
