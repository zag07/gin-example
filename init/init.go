package init

import (
	"github.com/zs368/gin-example/pkg/tracer"
	"log"

	"github.com/zs368/gin-example/configs"
	"github.com/zs368/gin-example/pkg/config"
	"github.com/zs368/gin-example/pkg/database"
	zLog "github.com/zs368/gin-example/pkg/log"
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

	if err := setTracer(); err != nil {
		log.Fatalf("init.setTracer err: %v", err)
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
	if zLog.Logger, err = zLog.NewLogger("storage/logs/logrus.log"); err != nil {
		return err
	}

	return nil
}

func setTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("gin-example", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	tracer.Tracer = jaegerTracer
	return nil
}
