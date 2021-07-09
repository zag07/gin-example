package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/zs368/gin-example/internal/pkg/app"
	"github.com/zs368/gin-example/internal/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := app.NewResponse(c)
		token := c.GetHeader("Authorization")

		if token == "" {
			r.ToErrorResponse(errcode.UnauthorizedTokenNotFound)
			c.Abort()
			return
		}

		claims, err := app.ParseToken(token)

		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				r.ToErrorResponse(errcode.UnauthorizedTokenTimeout)
			default:
				r.ToErrorResponse(errcode.UnauthorizedTokenError)
			}
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
