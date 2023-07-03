package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/retcode"
)

// 全局错误
func PathNotFound(ctx *gin.Context) {
	log.Errorf("request path:%s not found!", ctx.Request.RequestURI)
	app.ToResponseCode(ctx.Writer, retcode.RequestPathNotFound)
}

func MethodNotAllow(ctx *gin.Context) {
	log.Errorf("request path:%s method:%s not allowed!", ctx.Request.RequestURI, ctx.Request.Method)
	app.ToResponseCode(ctx.Writer, retcode.RequestMethodNotAllow)
}

// 中间件跳过
type SkipperFunc func(ctx *gin.Context) bool

func PathPrefixSkipper(prefixs ...string) SkipperFunc {
	return func(ctx *gin.Context) bool {
		path := ctx.Request.URL.Path
		pathLength := len(path)
		for _, prefix := range prefixs {
			pl := len(prefix)
			if pathLength >= pl && path[:pl] == prefix {
				return true
			}
		}
		return false
	}
}

func PathContainSkipper(prefixs ...string) SkipperFunc {
	return func(ctx *gin.Context) bool {
		path := ctx.Request.URL.Path
		for _, prefix := range prefixs {
			if strings.Contains(path, prefix) {
				return true
			}
		}
		return false
	}
}

func SkipHandler(c *gin.Context, skippers ...SkipperFunc) bool {
	for _, skipper := range skippers {
		if skipper(c) {
			return true
		}
	}
	return false
}
