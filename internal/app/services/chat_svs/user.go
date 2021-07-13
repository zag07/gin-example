package chat_svs

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

type User struct {
	broadcast *Broadcaster
	conn      *websocket.Conn
	send      chan []byte
}

// ServeWs handles websocket requests from the peer.
func ServeWs(b *Broadcaster, w http.ResponseWriter, r *http.Request) {
	coon, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	user := &User{
		broadcast: b,
		conn:      coon,
		send:      make(chan []byte, 256),
	}
	user.broadcast.register <- user

	go user.writePump()
	go user.readPump()
}

// readPump pumps messages from the websocket connection to the broadcast.
func (u *User) readPump() {
	defer func() {
		u.broadcast.unregister <- u
		u.conn.Close()
	}()

	u.conn.SetReadLimit(maxMessageSize)
	u.conn.SetReadDeadline(time.Now().Add(pongWait))
	u.conn.SetPongHandler(func(string) error {
		u.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := u.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		u.broadcast.messages <- message
	}
}

// writePump pumps messages from the broadcast to the websocket connection.
func (u *User) writePump() {
	tricker := time.NewTicker(pingPeriod)
	defer func() {
		tricker.Stop()
		u.conn.Close()
	}()

	for {
		select {
		case message, ok := <-u.send:
			u.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				u.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := u.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(u.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-u.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-tricker.C:
			u.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := u.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
