package configs

import "time"

var Server server

type server struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (s *server) Default() server {
	return server{
		RunMode:      "debug",
		HttpPort:     "8080",
		ReadTimeout:  60,
		WriteTimeout: 60,
	}
}
