package middleware

import (
	"context"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

const (
	// XRequestIDKey defines X-Request-ID key string.
	XRequestIDKey = "X-Request-ID"
)

func RequestId(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		rid := request.Header.Get(XRequestIDKey)
		if rid == "" {
			rid = uuid.NewV4().String()
		}
		ctx := context.WithValue(request.Context(), XRequestIDKey, rid)
		writer.Header().Set(XRequestIDKey, rid)
		handler.ServeHTTP(writer, request.WithContext(ctx))
	})
}
