package middleware

import (
	"blog-service/global"
	"context"
	"net/http"
)

// Timeout 超时控制
func Timeout(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx, cancel := context.WithTimeout(request.Context(), global.AppConfig.ContextTimeout)
		defer cancel()
		request = request.WithContext(ctx)
		handler.ServeHTTP(writer, request)
	})
}
