package middleware

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/retcode"
)

// JWT 身份验证
func JWT(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			app.ToResponseCode(c.Writer, retcode.RequestTokenEmpty)
			c.Abort()
		}
		token := strings.Split(auth, " ")
		if len(token) != 2 || token[0] != "bearer" {
			app.ToResponseCode(c.Writer, retcode.RequestTokenEmpty)
			c.Abort()
		}
		claims, err := app.ParseToken(token[1])
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				app.ToResponseCode(c.Writer, retcode.RequestTokenAuthExpire)
				c.Abort()
			default:
				app.ToResponseCode(c.Writer, retcode.RequestTokenAuthFail)
				c.Abort()
			}
		}
		c.Set("username", claims.UserName)
		c.Set("userid", claims.UserId)
		c.Next()
	}
}
