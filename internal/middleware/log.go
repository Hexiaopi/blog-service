package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/blog-service/global"
	"github.com/hexiaopi/blog-service/pkg/log"
)

type ResponseWithRecorder struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (rec *ResponseWithRecorder) WriteHeader(statusCode int) {
	rec.ResponseWriter.WriteHeader(statusCode)
	rec.statusCode = statusCode
}

func (rec *ResponseWithRecorder) Write(d []byte) (n int, err error) {
	n, err = rec.ResponseWriter.Write(d)
	if err != nil {
		return
	}
	rec.body.Write(d)
	return
}

// Logger 日志记录
func Logger(logger *log.Logger, skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}
		start := time.Now()
		//记录请求包
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr := ioutil.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = rdr //rewrite

		logger.Logger.Sugar().Infow("receive request",
			XRequestIDKey, c.GetString(XRequestIDKey),
			global.Path, c.Request.URL.Path,
			global.QueryParam, c.Request.URL.RawQuery,
			global.Method, c.Request.Method,
		)

		//记录返回包
		wc := &ResponseWithRecorder{
			ResponseWriter: c.Writer,
			statusCode:     http.StatusOK,
			body:           bytes.Buffer{},
		}

		c.Next()

		defer func() { //日志记录扫尾工作
			logger.Logger.Sugar().Infow("done",
				XRequestIDKey, c.GetString(XRequestIDKey),
				global.Path, c.Request.URL.Path,
				global.Status, wc.statusCode,
				global.ResPkg, wc.body.String(),
				"use time", time.Since(start).String(),
			)
		}()
	}
}
