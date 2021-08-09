package database

import (
	"fmt"

	"github.com/zs368/gin-example/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormOpentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

func SetDB() (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s",
		configs.MySQL.Username,
		configs.MySQL.Password,
		configs.MySQL.Host+":"+configs.MySQL.Port,
		configs.MySQL.DbName,
		configs.MySQL.Charset,
		true,
		"Local")
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = DB.Use(gormOpentracing.New()); err != nil {
		return nil, err
	}

	return DB, err
}
