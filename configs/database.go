package configs

import "github.com/zs368/gin-example/internal/pkg/config"

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
	ParseTime    bool
	Loc          string
	MaxIdleConns int
	MaxOpenConns int
}

func init() {
	Db.MySQL.Host = config.GetString("MYSQL_HOST", "127.0.0.1")

	Db.MySQL.Port = config.GetString("MYSQL_PORT", "33061")

	Db.MySQL.UserName = config.GetString("MYSQL_USERNAME", "root")

	Db.MySQL.Password = config.GetString("MYSQL_PASSWORD", "root")

	Db.MySQL.DBName = config.GetString("MYSQL_DATABASE", "gin_example")

	Db.MySQL.Charset = config.GetString("MYSQL_CHARSET", "utf8mb4")

	Db.MySQL.ParseTime = config.GetBool("MYSQL_PARSE_TIME", true)
}
