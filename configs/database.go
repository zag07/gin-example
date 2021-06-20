package configs

var Db database

type database struct {
	MySQL struct {
		UserName     string
		Password     string
		Host         string
		DBName       string
		TablePrefix  string
		Charset      string
		ParseTime    bool
		Loc          string
		MaxIdleConns int
		MaxOpenConns int
	}
}
