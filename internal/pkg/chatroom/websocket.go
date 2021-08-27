package chat_svs

import (
	"log"
	"net/http"
)

// ServeWs handles websocket requests from the peer.
func ServeWs(w http.ResponseWriter, r *http.Request) {
	coon, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// 1、构建新用户的实例
	user, err := NewUser(coon, r)
	if err != nil {
		log.Println(err)
		coon.Close()
		return
	}

	// 2、加入广播器的用户列表
	Broadcaster.entering <- user

	// 3、发送消息的 goroutine
	go user.writePump()

	// 4、接收消息的 goroutine
	go user.readPump()

	// 5、发送欢迎消息（只是这个用户）
	user.messages <- NewWelcomeMsg(user)

	// 6、像其余用户发送消息
	Broadcaster.Broadcast(NewUserEnterMsg(user))
}
