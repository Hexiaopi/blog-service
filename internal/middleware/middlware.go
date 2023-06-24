package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/retcode"
)

func PathNotFound(ctx *gin.Context) {
	log.Errorf("request path:%s not found!", ctx.Request.RequestURI)
	app.ToResponseCode(ctx.Writer, retcode.RequestPathNotFound)
}

func MethodNotAllow(ctx *gin.Context) {
	log.Errorf("request path:%s method:%s not allowed!", ctx.Request.RequestURI, ctx.Request.Method)
	app.ToResponseCode(ctx.Writer, retcode.RequestMethodNotAllow)
}

type SkipperFunc func(ctx *gin.Context) bool

func AllowPathPrefixShipper(prefixs ...string) SkipperFunc {
	return func(ctx *gin.Context) bool {
		path := ctx.Request.URL.Path
		pathLength := len(path)
		for _, p := range prefixs {
			pl := len(p)
			if pathLength >= pl && path[:pl] == p {
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
