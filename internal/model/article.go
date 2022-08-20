package model

import (
	"time"
)

type Article struct {
	ID            int       `gorm:"id"`
	Title         string    `gorm:"title"`
	Desc          string    `gorm:"desc"`
	Content       string    `gorm:"content"`
	CoverImageUrl string    `gorm:"cover_image_url"`
	State         uint8     `gorm:"state"`
	CreateTime    time.Time `gorm:"create_time"`
	UpdateTime    time.Time `gorm:"update_time"`
	Operator      string    `gorm:"operator"`
}

func (a Article) TableName() string {
	return "blog_article"
}
