package main

import (
	"log"

	"github.com/zs368/gin-example/configs"
	"github.com/zs368/gin-example/pkg/config"
	"github.com/zs368/gin-example/pkg/database"
	zLog "github.com/zs368/gin-example/pkg/log"
	"github.com/zs368/gin-example/pkg/trace"
)

func init() {
	if err := setConfig(); err != nil {
		log.Println("init.setConfig err: %v, %v", err, "将使用默认值")
	}

	if err := setTracer(); err != nil {
		log.Fatalf("init.setTracer err: %v", err)
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
	c.SetConfig("AUTH", configs.SetAuthConfig)
	c.SetConfig("DATABASE", configs.SetDbConfig)
	c.SetConfig("WS", configs.SetWSConfig)

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
	if err := zLog.InitLogger(); err != nil {
		return err
	}
	return nil
}

func setTracer() error {
	if _, err := trace.InitGlobalTracer("gin-example", "127.0.0.1:6831"); err != nil {
		return err
	}

	return nil
}
