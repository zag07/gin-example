package configs

import "github.com/zs368/gin-example/internal/pkg/config"

var App app

type app struct {
	port            string
	defaultPageSize int
	maxPageSize     int
}

func (a *app) Port() string {
	return config.GetString("APP_PORT", "8080")
}

func (a *app) DefaultPageSize() int {
	return config.GetInt("APP_DEFAULT_PAGE_SIZE", 30)
}

func (a *app) MaxPageSize() int {
	return config.GetInt("APP_MAX_PAGE_SIZE", 100)
}
