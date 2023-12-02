package entity

import (
	"time"

	"github.com/hexiaopi/blog-service/internal/model"
)

type Role struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	State      uint8  `json:"state"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func (r *Role) ToModel() *model.Role {
	createTime, _ := time.Parse(DefaultTimeFormat, r.CreateTime)
	updateTime, _ := time.Parse(DefaultTimeFormat, r.UpdateTime)
	return &model.Role{
		ID:         r.ID,
		Name:       r.Name,
		Desc:       r.Desc,
		State:      r.State,
		CreateTime: createTime,
		UpdateTime: updateTime,
	}
}

func ToEntityRole(role *model.Role) *Role {
	return &Role{
		ID:         role.ID,
		Name:       role.Name,
		Desc:       role.Desc,
		State:      role.State,
		CreateTime: role.CreateTime.Format(DefaultTimeFormat),
		UpdateTime: role.UpdateTime.Format(DefaultTimeFormat),
	}
}
