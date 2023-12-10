package entity

import (
	"time"

	"github.com/hexiaopi/blog-service/internal/model"
)

type OperationLog struct {
	ID         int    `json:"id"`
	UserAgent  string `json:"user_agent"`
	IP         string `json:"ip"`
	Object     string `json:"object"`
	Action     string `json:"action"`
	Result     string `json:"result"`
	Error      string `json:"error"`
	CreateTime string `json:"create_time"`
	User       User   `json:"user"`
}

func (l *OperationLog) ToModel() *model.OperationLog {
	createTime, _ := time.Parse(DefaultTimeFormat, l.CreateTime)
	log := model.OperationLog{
		ID:         l.ID,
		UserAgent:  l.UserAgent,
		IP:         l.IP,
		Object:     l.Object,
		Action:     l.Action,
		Result:     l.Result,
		Error:      l.Error,
		CreateTime: createTime,
		UserId:     l.User.ID,
	}
	return &log
}

func ToEntityOperation(log *model.OperationLog) *OperationLog {
	return &OperationLog{
		ID:         log.ID,
		UserAgent:  log.UserAgent,
		IP:         log.IP,
		Object:     log.Object,
		Action:     log.Action,
		Result:     log.Result,
		Error:      log.Error,
		CreateTime: log.CreateTime.Format(DefaultTimeFormat),
	}
}
