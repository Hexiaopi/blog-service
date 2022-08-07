package middleware

import (
	"net/http"
	"runtime/debug"

	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/retcode"
)

// Recovery 捕获异常，统一返回错误码
func Recovery(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic error %v", err)
				log.Printf(string(debug.Stack()))
				app.ToResponseCode(writer, retcode.UnknownError)
				return
			}
		}()
		handler.ServeHTTP(writer, request)
	})
}
