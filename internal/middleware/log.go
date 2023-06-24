package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/blog-service/global"
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
func Logger(skippers ...SkipperFunc) gin.HandlerFunc {
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

		log.WithFields(log.Fields{
			XRequestIDKey:     c.GetString(XRequestIDKey),
			global.Path:       c.Request.URL.Path,
			global.QueryParam: c.Request.URL.RawQuery,
			global.Method:     c.Request.Method,
		}).Infof("receive request body:%d ", len(buf))

		//记录返回包
		wc := &ResponseWithRecorder{
			ResponseWriter: c.Writer,
			statusCode:     http.StatusOK,
			body:           bytes.Buffer{},
		}

		c.Next()

		defer func() { //日志记录扫尾工作
			log.WithFields(log.Fields{
				XRequestIDKey: c.GetString(XRequestIDKey),
				global.Path:   c.Request.URL.Path,
				global.Status: wc.statusCode,
				global.ResPkg: wc.body.String(),
			}).Infof("done use time:%s", time.Since(start).String())
		}()
	}
}
