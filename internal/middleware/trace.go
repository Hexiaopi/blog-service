package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"github.com/hexiaopi/blog-service/internal/config"
)

// Tracer 调用链追踪
func Tracer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx context.Context
		var span opentracing.Span
		spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			span, ctx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				config.Tracer,
				c.Request.URL.Path,
				opentracing.Tag{Key: XRequestIDKey, Value: c.GetString(XRequestIDKey)})
		} else {
			span, ctx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				config.Tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"})
		}
		span.SetTag("http.method", c.Request.Method)
		span.SetTag("http.url", c.Request.URL.String())
		defer span.Finish()
		ctxSpan := opentracing.ContextWithSpan(ctx, span)
		c.Request = c.Request.WithContext(ctxSpan)
		c.Next()
	}
}
