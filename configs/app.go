package configs

import "github.com/zs368/gin-example/pkg/config"

var App app

type app struct {
	Port                 string
	DefaultPageSize      int
	MaxPageSize          int
	UploadSavePath       string
	UploadServerUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
}

func SetAppConfig(c *config.Config) {
	App.Port = c.GetString("APP_PORT", "8080")

	App.DefaultPageSize = c.GetInt("APP_DEFAULT_PAGE_SIZE", 30)

	App.MaxPageSize = c.GetInt("APP_MAX_PAGE_SIZE", 100)

	App.UploadSavePath = c.GetString("UploadSavePath", "storage/app/uploads")

	App.UploadServerUrl = c.GetString("UploadServerUrl", "http://127.0.0.1:8080/static")

	App.UploadImageMaxSize = c.GetInt("UploadImageMaxSize", 5)

	App.UploadImageAllowExts = c.GetStringSlice("UploadImageAllowExts", ".jpg,.jpeg,.png")
}
