package middleware

import (
	"context"
	"net/http"

	"github.com/hexiaopi/blog-service/internal/config"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// Tracer 调用链追踪
func Tracer(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var ctx context.Context
		var span opentracing.Span
		spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(request.Header))
		if err != nil {
			span, ctx = opentracing.StartSpanFromContextWithTracer(
				request.Context(),
				config.Tracer,
				request.URL.Path,
				opentracing.Tag{Key: XRequestIDKey, Value: request.Context().Value(XRequestIDKey)})
		} else {
			span, ctx = opentracing.StartSpanFromContextWithTracer(
				request.Context(),
				config.Tracer,
				request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"})
		}
		span.SetTag("http.method", request.Method)
		span.SetTag("http.url", request.URL.String())
		defer span.Finish()
		ctxSpan := opentracing.ContextWithSpan(ctx, span)
		handler.ServeHTTP(writer, request.WithContext(ctxSpan))
	})
}
