package chat_svs

import "container/ring"

type offlineProcessor struct {
	n          int
	recentRing *ring.Ring
}

var n = 10

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
