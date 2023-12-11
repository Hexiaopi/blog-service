package entity

import (
	"github.com/hexiaopi/blog-service/internal/model"
)

type SysRest struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Path       string    `json:"path"`
	Method     string    `json:"method"`
	CreateTime string    `json:"create_time"`
	UpdateTime string    `json:"update_time"`
	ParentId   int       `json:"parent_id"`
	Children   []SysRest `json:"children"`
}

func ToEntitySysRest(sysRest *model.SysRest) *SysRest {
	result := SysRest{
		ID:         sysRest.ID,
		Name:       sysRest.Name,
		Path:       sysRest.Path,
		Method:     sysRest.Method,
		ParentId:   sysRest.ParentId,
		CreateTime: sysRest.CreateTime.Format(DefaultTimeFormat),
		UpdateTime: sysRest.UpdateTime.Format(DefaultTimeFormat),
	}
	return &result
}
