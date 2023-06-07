package middleware

import (
	"context"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/blog-service/internal/app"
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
func (op *Operation) RecordOperation(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			path := request.URL.Path
			object := strings.TrimPrefix(path, "/api/v1/")

			handler.ServeHTTP(writer, request)

			go func() {
				operation := service.CreateOperationRequest{
					SystemOperationLog: model.SystemOperationLog{
						UserId:    request.Context().Value("userid").(int),
						UserAgent: request.UserAgent(),
						IP:        app.GetRemoteIp(request),
						Object:    strings.Split(object, "/")[0],
						Action:    request.Method,
						Result:    "",
					},
				}
				if err := op.srv.Operations().Create(context.Background(), &operation); err != nil {
					log.Errorf("record operation log err:%v", err)
				}
			}()
		} else {
			handler.ServeHTTP(writer, request)
		}
	})
}
