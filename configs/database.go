package configs

import "github.com/zs368/gin-example/pkg/config"

var Db db

type db struct {
	MySQL mysql
}

type mysql struct {
	Host         string
	Port         string
	UserName     string
	Password     string
	DBName       string
	TablePrefix  string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
}

func SetDbConfig(c *config.Config) {
	Db.MySQL.Host = c.GetString("MYSQL_HOST", "127.0.0.1")

	Db.MySQL.Port = c.GetString("MYSQL_PORT", "33061")

	Db.MySQL.UserName = c.GetString("MYSQL_USERNAME", "root")

	Db.MySQL.Password = c.GetString("MYSQL_PASSWORD", "root")

	Db.MySQL.DBName = c.GetString("MYSQL_DATABASE", "gin_example")

	Db.MySQL.Charset = c.GetString("MYSQL_CHARSET", "utf8mb4")
}
