package model

import "time"

type SystemOperationLog struct {
	ID         int       `gorm:"primaryKey;autoIncrement;id" json:"id"`
	UserId     int       `gorm:"user_id" json:"user_id"`
	UserAgent  string    `gorm:"user_agent" json:"user_agent"`
	IP         string    `gorm:"ip" json:"ip"`
	Object     string    `gorm:"object" json:"object"`
	Action     string    `gorm:"action" json:"action"`
	Result     string    `gorm:"result" json:"result"`
	CreateTime time.Time `gorm:"create_time" json:"create_time,omitempty"`
}

func (so SystemOperationLog) TableName() string {
	return "system_operation_log"
}
