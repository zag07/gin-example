package chat_svs

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zs368/gin-example/internal/pkg/app"
)

var (
	writeWait      time.Duration = 10
	pongWait       time.Duration = 60
	pingPeriod                   = (pongWait * 9) / 10
	maxMessageSize               = int64(512)
	newline                      = []byte{'\n'}
	space                        = []byte{' '}

	globalUID uint32 = 0
)

var upgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

type User struct {
	Uid      uint      `json:"uid"`
	Username string    `json:"nickname"`
	EnterAt  time.Time `json:"enter_at"`
	Addr     string    `json:"addr"`
	Token    string    `json:"token"`

	conn     *websocket.Conn
	send     chan []byte
	messages chan *Message
}

var system = &User{}

func NewUser(coon *websocket.Conn, r *http.Request) (*User, error) {
	user := &User{
		Username: r.FormValue("nickname"),
		EnterAt:  time.Now(),
		Addr:     r.RemoteAddr,

		conn:     coon,
		send:     make(chan []byte, 256),
		messages: make(chan *Message, 32),
	}

	if user.Token != "" {
		claims, err := app.ParseToken(r.FormValue("token"))
		if err != nil {
			user.messages <- NewErrorMsg("token 生成失败")
			return nil, err
		}
		user.Uid = claims.Uid
	}

	if user.Uid == 0 {
		user.Uid = uint(atomic.AddUint32(&globalUID, 1))
		token, err := app.GenerateToken(app.UserInfo{Uid: user.Uid, Username: user.Username})
		if err != nil {
			user.messages <- NewErrorMsg("token 解析失败")
			return nil, err
		}
		user.Token = token
	}

	return user, nil
}

// readPump pumps messages from the websocket connection to the broadcaster.
func (u *User) readPump() {
	defer func() {
		Broadcaster.leaving <- u
		Broadcaster.Broadcast(NewUserLeaveMsg(u))
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

		if bytes.Index(message, []byte("content")) != -1 {
			var msg map[string]string
			if err := json.Unmarshal(message, &msg); err != nil {
				log.Printf("error: %v", err)
				break
			}

			Broadcaster.messages <- NewMsg(u, msg["content"])
		} else {
			Broadcaster.send <- message
		}
	}
}

// writePump pumps messages from the broadcaster to the websocket connection.
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
		case message, ok := <-u.messages:
			u.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				u.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			u.conn.WriteJSON(message)
		case <-tricker.C:
			u.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := u.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
