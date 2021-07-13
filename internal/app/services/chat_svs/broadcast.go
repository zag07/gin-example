package chat_svs

type Broadcaster struct {
	users      map[*User]bool
	register   chan *User
	unregister chan *User
	messages   chan []byte
}

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{
		users:      make(map[*User]bool),
		register:   make(chan *User),
		unregister: make(chan *User),
		messages:   make(chan []byte),
	}
}

func (b *Broadcaster) Run() {
	for {
		select {
		case user := <-b.register:
			b.users[user] = true
		case user := <-b.unregister:
			if _, ok := b.users[user]; ok {
				delete(b.users, user)
				close(user.send)
			}
		case message := <-b.messages:
			for user := range b.users {
				select {
				case user.send <- message:
				default:
					delete(b.users, user)
					close(user.send)
				}
			}

		}
	}
}
