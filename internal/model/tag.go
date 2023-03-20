package model

import (
	"time"
)

type Tag struct {
	ID           int       `gorm:"id" json:"id"`
	Name         string    `gorm:"name" json:"name"`
	Desc         string    `gorm:"desc" json:"desc"`
	State        uint8     `gorm:"state" json:"state"`
	CreateTime   time.Time `gorm:"create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"update_time" json:"update_time"`
	Operator     string    `gorm:"operator" json:"operator"`
	ArticleTotal int       `gorm:"-" json:"article_total"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
