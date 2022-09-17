package middleware

import (
	"context"
	"net/http"

	"github.com/hexiaopi/blog-service/internal/config"
)

// Timeout 超时控制
func Timeout(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx, cancel := context.WithTimeout(request.Context(), config.AppEngine.ContextTimeout)
		defer cancel()
		request = request.WithContext(ctx)
		handler.ServeHTTP(writer, request)
	})
}
