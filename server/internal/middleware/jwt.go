package middleware

import (
	"blog-service/internal/app"
	"blog-service/internal/retcode"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

// JWT 身份验证
func JWT(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		code := retcode.Success
		token := request.Header.Get("X-Token")
		if token == "" {
			code = retcode.RequestTokenEmpty
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = retcode.RequestTokenAuthTimeout
				default:
					code = retcode.RequestTokenAuthFail
				}
			}
		}
		if code != retcode.Success {
			app.ToResponseCode(writer, code)
			return
		}
		handler.ServeHTTP(writer, request)
	})
}
