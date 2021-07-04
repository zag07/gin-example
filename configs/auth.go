package configs

import "github.com/zs368/gin-example/pkg/config"

var Auth auth

type auth struct {
	JwtSecret string
	JwtIssuer string
	JwtExpire uint
}

func SetAuthConfig(c *config.Config) {
	Auth.JwtSecret = c.GetString("JWT_SECRET", "echo")

	Auth.JwtIssuer = c.GetString("JWT_ISSUER", "gin-example")

	Auth.JwtExpire = c.GetUint("JWT_EXPIRE", 7200)
}
