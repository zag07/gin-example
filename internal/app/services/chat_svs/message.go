package chat_svs

import "time"

const (
	MsgTypeNormal = iota
	MsgTypeError
	MsgTypeWelcome
	MsgTypeUserEnter
	MsgTypeUserLeave
)

type Message struct {
	User    *User     `json:"user"`
	Type    int       `json:"type"`
	Content string    `json:"content"`
	MsgTime time.Time `json:"msg_time"`
}

func NewMsg(u *User, content string) *Message {
	return &Message{
		User:    u,
		Type:    MsgTypeNormal,
		Content: content,
		MsgTime: time.Time{},
	}
}

func NewErrorMsg(content string) *Message {
	return &Message{
		User:    system,
		Type:    MsgTypeError,
		Content: content,
		MsgTime: time.Now(),
	}
}

func NewWelcomeMsg(u *User) *Message {
	return &Message{
		User:    u,
		Type:    MsgTypeWelcome,
		Content: u.Username + " 您好，欢迎加入聊天室！",
		MsgTime: time.Now(),
	}
}

func NewUserEnterMsg(u *User) *Message {
	return &Message{
		User:    u,
		Type:    MsgTypeUserEnter,
		Content: u.Username + " 加入了聊天室",
		MsgTime: time.Now(),
	}
}

func NewUserLeaveMsg(u *User) *Message {
	return &Message{
		User:    u,
		Type:    MsgTypeUserLeave,
		Content: u.Username + " 离开了聊天室",
		MsgTime: time.Now(),
	}
}
