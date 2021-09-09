package data

import (
	"gorm.io/gorm"

	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	// "github.com/google/wire"
	"go.uber.org/zap"
	// init mysql driver
	"gorm.io/driver/mysql"

	"github.com/zag07/gin-example/internal/conf"
)

// ProviderSet is data providers.
// var ProviderSet = wire.NewSet(NewData, NewBlogRepo)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewData(conf *conf.Data, log *zap.Logger) (*Data, func(), error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conf.Database.Source, // DSN data source name
		DefaultStringSize:         256,                  // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                 // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                 // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                 // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		log.Sugar().Errorf("err:%v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	d := &Data{
		db:  db,
		rdb: rdb,
	}

	return d, func() {
		log.Info("message closing the data resources")
		sqlDB, err := db.DB()
		if err != nil {
			log.Sugar().Errorf("err:%v", err)
		}
		if err := sqlDB.Close(); err != nil {
			log.Sugar().Errorf("err:%v", err)
		}
		if err := d.rdb.Close(); err != nil {
			log.Sugar().Errorf("err:%v", err)
		}
	}, nil
}
