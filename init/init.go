package init

import (
	"log"

	"github.com/zs368/gin-example/configs"
	"github.com/zs368/gin-example/internal/pkg/config"
)

func init() {
	if err := config.GetConfig("App", &configs.App); err != nil {
		log.Fatalf("config.GetConfig err: %v", err)
	}
	if err := config.GetConfig("Database", &configs.Db); err != nil {
		log.Fatalf("config.GetConfig err: %v", err)
	}
}
