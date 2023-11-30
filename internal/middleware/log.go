package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/blog-service/global"
	log "github.com/hexiaopi/blog-service/pkg/logger"
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
func Logger(logger log.Logger, skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}
		start := time.Now()
		//记录请求包
		buf, _ := io.ReadAll(c.Request.Body)
		rdr := io.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = rdr //rewrite

		logger.Info("receive request",
			log.String(XRequestIDKey, c.GetString(XRequestIDKey)),
			log.String(global.Path, c.Request.URL.Path),
			log.String(global.QueryParam, c.Request.URL.RawQuery),
			log.String(global.Method, c.Request.Method),
		)

		//记录返回包
		wc := &ResponseWithRecorder{
			ResponseWriter: c.Writer,
			statusCode:     http.StatusOK,
			body:           bytes.Buffer{},
		}

		c.Next()

		defer func() { //日志记录扫尾工作
			logger.Info("done request",
				log.String(XRequestIDKey, c.GetString(XRequestIDKey)),
				log.String(global.Path, c.Request.URL.Path),
				log.Int(global.Status, wc.statusCode),
				log.String(global.ResPkg, wc.body.String()),
				log.String(global.UseTime, time.Since(start).String()),
			)
		}()
	}
}
