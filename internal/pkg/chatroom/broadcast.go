package chat_svs

import (
	"log"
)

type broadcaster struct {
	users    map[*User]bool
	entering chan *User
	leaving  chan *User
	send     chan []byte
	messages chan *Message
}


var Broadcaster = &broadcaster{
	users:    make(map[*User]bool),
	entering: make(chan *User),
	leaving:  make(chan *User),
	send:     make(chan []byte),
	messages: make(chan *Message, 1000),
}

func (b *broadcaster) Run() {
	for {
		select {
		case user := <-b.entering:
			b.users[user] = true
			OfflineProcessor.Send(user)
		case user := <-b.leaving:
			if _, ok := b.users[user]; ok {
				delete(b.users, user)
				close(user.send)
			}
		case s := <-b.send:
			for user := range b.users {
				select {
				case user.send <- s:
				default:
					delete(b.users, user)
					close(user.send)
				}
			}
		case msg := <-b.messages:
			for user := range b.users {
				if user.Uid == msg.User.Uid {
					continue
				}
				select {
				case user.messages <- msg:
				default:
					delete(b.users, user)
					close(user.send)
				}
			}
			OfflineProcessor.Save(msg)
		}
	}
}

func (b *broadcaster) Broadcast(msg *Message) {
	if len(b.messages) >= int(1000) {
		log.Println("broadcast queue 满了")
	}
	b.messages <- msg
}

func (b *broadcaster) GetUserList() []*User {
	userList := make([]*User, 0, len(b.users))
	for user, exist := range b.users {
		if exist == true {
			userList = append(userList, user)
		}
	}
	return userList
}
