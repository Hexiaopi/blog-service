package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hexiaopi/blog-service/internal/config"
)

// Timeout 超时控制
func Timeout() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), config.AppEngine.ContextTimeout)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
