package middleware

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/retcode"
)

// JWT 身份验证
func JWT(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("X-Token")
		if token == "" {
			app.ToResponseCode(writer, retcode.RequestTokenEmpty)
			return
		}
		claims, err := app.ParseToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				app.ToResponseCode(writer, retcode.RequestTokenAuthTimeout)
				return
			default:
				app.ToResponseCode(writer, retcode.RequestTokenAuthFail)
				return
			}
		}
		ctx := context.WithValue(request.Context(), "name", claims.UserName)
		request = request.WithContext(ctx)
		handler.ServeHTTP(writer, request)
	})
}
