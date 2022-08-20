package model

import (
	"time"
)

type Tag struct {
	ID         int       `gorm:"id"`
	Name       string    `gorm:"name"`
	Desc       string    `gorm:"desc"`
	State      uint8     `gorm:"state"`
	CreateTime time.Time `gorm:"create_time"`
	UpdateTime time.Time `gorm:"update_time"`
	Operator   string    `gorm:"operator"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
