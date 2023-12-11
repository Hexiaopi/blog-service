package entity

import "strings"

type OneOption struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	State    uint8  `json:"state"`
	Operator string `json:"operator"`
}

type ListOption struct {
	Name     string `json:"name"`
	UserId   int    `json:"user_id"`
	Object   string `json:"object"`
	Action   string `json:"action"`
	Result   string `json:"result"`
	Sort     string `json:"sort"`
	State    uint8  `json:"state"`
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	ParentId int    `json:"parent_id"`
}

func (option ListOption) GetPageOffset() int {
	return (option.Page - 1) * option.Limit
}

func (option ListOption) GetSortType() string {
	sortType := ""
	if strings.HasPrefix(option.Sort, "+") {
		sortType = strings.TrimPrefix(option.Sort, "+") + " asc"
	}
	if strings.HasPrefix(option.Sort, "-") {
		sortType = strings.TrimPrefix(option.Sort, "-") + " desc"
	}
	return sortType
}

const (
	DefaultTimeFormat = "2006-01-02 15:04:05"
)
