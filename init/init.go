package init

import (
	"log"

	"github.com/zs368/gin-example/configs"
	"github.com/zs368/gin-example/internal/pkg/config"
)

func init() {
	if err := setConfig(); err != nil {
		log.Fatalf("init.setConfig err: %v", err)
	}
}

func setConfig() error {
	c, err := config.NewConfig(".env")
	if err != nil {
		return err
	}

	c.SetConfig("APP", configs.SetAppConfig)
	c.SetConfig("DATABASE", configs.SetDbConfig)

	return nil
}
