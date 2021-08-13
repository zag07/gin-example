package chat_svs

import (
	"container/ring"

	"github.com/zs368/gin-example/configs"
)

type offlineProcessor struct {
	n          int
	recentRing *ring.Ring
}

var n = int(configs.App.WsOfflineNum)

var OfflineProcessor = &offlineProcessor{
	n:          n,
	recentRing: ring.New(n),
}

func (o *offlineProcessor) Save(msg *Message) {
	if msg.Type != MsgTypeNormal {
		return
	}
	o.recentRing.Value = msg
	o.recentRing = o.recentRing.Next()
}

func (o *offlineProcessor) Send(u *User) {
	o.recentRing.Do(func(i interface{}) {
		if i != nil {
			u.messages <- i.(*Message)
		}
	})
}
