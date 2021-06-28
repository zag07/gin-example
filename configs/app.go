package configs

import "github.com/zs368/gin-example/pkg/config"

var App app

type app struct {
	Port            string
	DefaultPageSize int
	MaxPageSize     int
}

func SetAppConfig(c *config.Config) {
	App.Port = c.GetString("APP_PORT", "8080")

	App.DefaultPageSize = c.GetInt("APP_DEFAULT_PAGE_SIZE", 30)

	App.MaxPageSize = c.GetInt("APP_MAX_PAGE_SIZE", 100)
}
