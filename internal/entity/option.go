package entity

import (
	"github.com/hexiaopi/blog-service/internal/app"
)

type GetOption struct {
	Id    int   `json:"id"`
	State uint8 `json:"state"`
}

type ListOption struct {
	Name  string `json:"name"`
	Sort  string `json:"sort"`
	State uint8  `json:"state"`
	*app.Page
}
