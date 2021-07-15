package configs

import (
	"time"

	"github.com/zs368/gin-example/pkg/config"
)

var WS ws

type ws struct {
	WriteWait      time.Duration
	PongWait       time.Duration
	MaxMessageSize int64
	MessageQueue   int
	OfflineNum     int
}

func SetWSConfig(c *config.Config) {
	WS.WriteWait = c.GetDuration("WS_Write_Wait", 10*time.Second)

	WS.PongWait = c.GetDuration("WS_Pong_Wait", 60*time.Second)

	WS.MaxMessageSize = c.GetInt64("WS_Max_Message_Size", 512)

	WS.MessageQueue = c.GetInt("WS_Message_Queue", 1024)

	WS.OfflineNum = c.GetInt("WS_Offline_Num", 10)
}
