package model

import (
	"time"
)

type Article struct {
	ID         int       `gorm:"primaryKey;autoIncrement;id" json:"id"`
	Name       string    `gorm:"name" json:"name"`
	Desc       string    `gorm:"desc" json:"desc"`
	Content    string    `gorm:"content" json:"content"`
	State      uint8     `gorm:"state" json:"state"`
	CreateTime time.Time `gorm:"create_time" json:"create_time,omitempty"`
	UpdateTime time.Time `gorm:"update_time" json:"update_time,omitempty"`
	Operator   string    `gorm:"operator" json:"operator"`
	Tags       []Tag     `gorm:"many2many:blog_article_tag" json:"tags"`
}

func (a Article) TableName() string {
	return "blog_article"
}

// 连接表：blog_article_tag
//   foreign key: article_id, reference: articles.id
//   foreign key: tag_id, reference: tags.id
