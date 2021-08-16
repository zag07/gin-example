package data

import (
	"github.com/zs368/gin-example/internal/conf"
	"github.com/zs368/gin-example/internal/data/ent"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Data .
type Data struct {
	db  *ent.Client
	rdb *redis.Client
}

func NewData(conf *conf.Data, logger zap.Logger) (*Data, func(), error) {

	return nil, nil, nil
}
