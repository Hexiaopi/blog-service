package model

import "time"

type OperationLog struct {
	ID         int       `gorm:"primaryKey;autoIncrement;id" json:"id"`
	UserId     int       `gorm:"user_id" json:"user_id"`
	User       User      `json:"user"`
	UserAgent  string    `gorm:"user_agent" json:"user_agent"`
	IP         string    `gorm:"ip" json:"ip"`
	Object     string    `gorm:"object" json:"object"`
	Action     string    `gorm:"action" json:"action"`
	Result     string    `gorm:"result" json:"result"`
	CreateTime time.Time `gorm:"create_time" json:"create_time,omitempty"`
}

func (so OperationLog) TableName() string {
	return "sys_operation_log"
}
