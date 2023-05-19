package model

import "time"

type SystemConfig struct {
	ID         int       `gorm:"primaryKey;autoIncrement;id" json:"id"`
	Name       string    `gorm:"name" json:"name"`
	Desc       string    `gorm:"desc" json:"desc"`
	Value      string    `gorm:"value" json:"value"`
	CreateTime time.Time `gorm:"create_time" json:"create_time,omitempty"`
	UpdateTime time.Time `gorm:"update_time" json:"update_time,omitempty"`
	Operator   string    `gorm:"operator" json:"operator"`
}

func (sc SystemConfig) TableName() string {
	return "system_config"
}
