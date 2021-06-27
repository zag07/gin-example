package init

import (
	"log"

	"github.com/zs368/gin-example/configs"
	"github.com/zs368/gin-example/internal/pkg/config"
	"github.com/zs368/gin-example/internal/pkg/database"
	"github.com/zs368/gin-example/internal/pkg/logger"
)

func init() {
	if err := setConfig(); err != nil {
		log.Fatalf("init.setConfig err: %v", err)
	}

	if err := setDatabase(); err != nil {
		log.Fatalf("init.setDatabase err: %v", err)
	}

	if err := setLogger(); err != nil {
		log.Fatalf("init.setLogger err: %v", err)
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

func setDatabase() error {
	var err error
	if database.DB, err = database.NewDB(); err != nil {
		return err
	}

	return nil
}

func setLogger() error {
	var err error
	if logger.Log, err = logger.NewLogger("storage/logs/logrus.log"); err != nil {
		return err
	}

	return nil
}
