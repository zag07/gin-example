package init

import (
	"log"

	"github.com/zs368/gin-example/configs"
	"github.com/zs368/gin-example/pkg/config"
	"github.com/zs368/gin-example/pkg/database"
	"github.com/zs368/gin-example/pkg/logger"
)

func init() {
	if err := setConfig(); err != nil {
		log.Println("init.setConfig err: %v, %v", err, "将使用默认值")
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
