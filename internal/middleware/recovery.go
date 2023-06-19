package middleware

import (
	"runtime/debug"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/retcode"
)

// Recovery 捕获异常，统一返回错误码
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic error %v", err)
				log.Printf(string(debug.Stack()))
				app.ToResponseCode(c.Writer, retcode.UnknownError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
