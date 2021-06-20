package configs

import "time"

var App app

type app struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (a app) Default() app {
	return app{
		RunMode:      "debug",
		HttpPort:     "8080",
		ReadTimeout:  60,
		WriteTimeout: 60,
	}
}
