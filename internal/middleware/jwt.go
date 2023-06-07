package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/retcode"
)

// JWT 身份验证
func JWT(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		auth := request.Header.Get("Authorization")
		if auth == "" {
			app.ToResponseCode(writer, retcode.RequestTokenEmpty)
			return
		}
		token := strings.Split(auth, " ")
		if len(token) != 2 || token[0] != "bearer" {
			app.ToResponseCode(writer, retcode.RequestTokenEmpty)
			return
		}
		claims, err := app.ParseToken(token[1])
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				app.ToResponseCode(writer, retcode.RequestTokenAuthExpire)
				return
			default:
				app.ToResponseCode(writer, retcode.RequestTokenAuthFail)
				return
			}
		}
		ctx := context.WithValue(request.Context(), "name", claims.UserName)
		ctx = context.WithValue(ctx, "userid", claims.UserId)
		request = request.WithContext(ctx)
		handler.ServeHTTP(writer, request)
	})
}
