package entity

import "github.com/hexiaopi/blog-service/internal/model"

type Config struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Value      string `json:"value"`
	CreateTime string `json:"create_time,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`
	Operator   string `json:"operator"`
}

func ToEntityConfig(config *model.Config) *Config {
	result := Config{
		ID:         config.ID,
		Name:       config.Name,
		Value:      config.Value,
		CreateTime: config.CreateTime.Format(DefaultTimeFormat),
		UpdateTime: config.UpdateTime.Format(DefaultTimeFormat),
		Operator:   config.Operator,
	}
	return &result
}
