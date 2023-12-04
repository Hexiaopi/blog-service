package entity

import (
	"time"

	"github.com/hexiaopi/blog-service/internal/model"
)

type Tag struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	State        uint8  `json:"state"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
	Operator     string `json:"operator"`
	ArticleTotal int    `json:"article_total"`
}

func (t *Tag) ToModel() *model.Tag {
	createTime, _ := time.Parse(DefaultTimeFormat, t.CreateTime)
	updateTime, _ := time.Parse(DefaultTimeFormat, t.UpdateTime)
	return &model.Tag{
		ID:         t.ID,
		Name:       t.Name,
		Desc:       t.Desc,
		State:      t.State,
		CreateTime: createTime,
		UpdateTime: updateTime,
		Operator:   t.Operator,
	}
}

func ToEntityTag(tag *model.Tag) *Tag {
	return &Tag{
		ID:         tag.ID,
		Name:       tag.Name,
		Desc:       tag.Desc,
		State:      tag.State,
		CreateTime: tag.CreateTime.Format(DefaultTimeFormat),
		UpdateTime: tag.UpdateTime.Format(DefaultTimeFormat),
		Operator:   tag.Operator,
	}
}
