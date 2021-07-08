package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/zs368/gin-example/internal/pkg/app"
	"github.com/zs368/gin-example/internal/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			eCode = errcode.Success
			token = c.GetHeader("Authorization")
		)

		if token == "" {
			eCode = errcode.InvalidParams.WithDetails("Authorization 未找到")
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					eCode = errcode.UnauthorizedTokenTimeout
				default:
					eCode = errcode.UnauthorizedTokenError
				}
			}
		}

		if eCode != errcode.Success {
			app.NewResponse(c).ToErrorResponse(eCode)
			c.Abort()
			return
		}

		c.Next()
	}
}
