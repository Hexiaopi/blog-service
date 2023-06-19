package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/service"
	"github.com/hexiaopi/blog-service/internal/store"
)

type Operation struct {
	srv service.Service
}

func NewOperation(store store.Factory) *Operation {
	return &Operation{
		srv: service.NewService(store, nil),
	}
}

// Logger 日志记录
func (op *Operation) RecordOperation() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodGet {
			path := c.Request.URL.Path
			object := strings.TrimPrefix(path, "/api/v1/")

			c.Next()

			go func() {
				operation := service.CreateOperationRequest{
					SystemOperationLog: model.SystemOperationLog{
						UserId:    c.GetInt("userid"),
						UserAgent: c.Request.UserAgent(),
						IP:        c.RemoteIP(),
						Object:    strings.Split(object, "/")[0],
						Action:    c.Request.Method,
						Result:    "",
					},
				}
				if err := op.srv.Operations().Create(context.Background(), &operation); err != nil {
					log.Errorf("record operation log err:%v", err)
				}
			}()
		} else {
			c.Next()
		}
	}
}
