package entity

import (
	"time"

	"github.com/hexiaopi/blog-service/internal/model"
)

type Resource struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Blob       []byte `json:"blob"`
	Base64     string `json:"base64"`
	Type       string `json:"type"`
	Size       int64  `json:"size"`
	State      uint8  `json:"state"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	Operator   string `json:"operator"`
}

func (r *Resource) ToModel() *model.Resource {
	createTime, _ := time.Parse(DefaultTimeFormat, r.CreateTime)
	updateTime, _ := time.Parse(DefaultTimeFormat, r.UpdateTime)
	return &model.Resource{
		ID:         r.ID,
		Name:       r.Name,
		Blob:       r.Blob,
		Base64:     r.Base64,
		Type:       r.Type,
		Size:       r.Size,
		State:      r.State,
		CreateTime: createTime,
		UpdateTime: updateTime,
		Operator:   r.Operator,
	}
}

func ToEntityResource(resource *model.Resource) *Resource {
	return &Resource{
		ID:         resource.ID,
		Name:       resource.Name,
		Blob:       resource.Blob,
		Base64:     resource.Base64,
		Type:       resource.Type,
		Size:       resource.Size,
		State:      resource.State,
		CreateTime: resource.CreateTime.Format(DefaultTimeFormat),
		UpdateTime: resource.UpdateTime.Format(DefaultTimeFormat),
		Operator:   resource.Operator,
	}
}
