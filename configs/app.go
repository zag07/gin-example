package configs

var App app

type app struct {
	DefaultPageSize int
	MaxPageSize     int
}

func (a *app) Default() app {
	return app{
		DefaultPageSize: 30,
		MaxPageSize:     100,
	}
}
