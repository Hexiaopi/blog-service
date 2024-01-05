package entity

import "github.com/hexiaopi/blog-service/internal/model"

type MenuTree struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Path       string     `json:"path"`
	Meta       MenuMeta   `json:"meta"`
	Component  string     `json:"component"`
	Redirect   string     `json:"redirect"`
	Hidden     bool       `json:"hidden"`
	Order      int        `json:"order"`
	ParentId   int        `json:"parent_id"`
	CreateTime string     `json:"create_time"`
	UpdateTime string     `json:"update_time"`
	Children   []MenuTree `json:"children,omitempty"`
}

type MenuMeta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

func ToEntityMenuTree(menu *model.SysMenu) *MenuTree {
	return &MenuTree{
		ID:   menu.ID,
		Name: menu.Name,
		Path: menu.Path,
		Meta: MenuMeta{
			Title: menu.Title,
			Icon:  menu.Icon,
		},
		Component:  menu.Component,
		Redirect:   menu.Redirect,
		Hidden:     menu.Hidden,
		Order:      menu.Sort,
		CreateTime: menu.CreateTime.Format(DefaultTimeFormat),
		UpdateTime: menu.UpdateTime.Format(DefaultTimeFormat),
	}
}

type SysMenu struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Path       string    `json:"path"`
	Title      string    `json:"title"`
	Icon       string    `json:"icon"`
	Component  string    `json:"component"`
	Redirect   string    `json:"redirect"`
	Hidden     bool      `json:"hidden"`
	Order      int       `json:"order"`
	ParentId   int       `json:"parent_id"`
	CreateTime string    `json:"create_time"`
	UpdateTime string    `json:"update_time"`
	Children   []SysMenu `json:"children,omitempty"`
}

func ToEntitySysMenu(menu *model.SysMenu) *SysMenu {
	return &SysMenu{
		ID:         menu.ID,
		Name:       menu.Name,
		Path:       menu.Path,
		Title:      menu.Title,
		Icon:       menu.Icon,
		Component:  menu.Component,
		Redirect:   menu.Redirect,
		Hidden:     menu.Hidden,
		Order:      menu.Sort,
		CreateTime: menu.CreateTime.Format(DefaultTimeFormat),
		UpdateTime: menu.UpdateTime.Format(DefaultTimeFormat),
	}
}
