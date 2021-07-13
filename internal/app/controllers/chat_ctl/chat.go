package chat_ctl

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/zs368/gin-example/internal/app/services/chat_svs"
)

type Chat struct{}

func NewChat() Chat {
	return Chat{}
}

func (c Chat) Home(ctx *gin.Context) {
	var (
		w = ctx.Writer
		r = ctx.Request
	)

	http.ServeFile(w, r, "./assets/template/home.html")
}

func (c Chat) WS(b *chat_svs.Broadcaster) gin.HandlerFunc {
	return func(c *gin.Context) {
		chat_svs.ServeWs(b, c.Writer, c.Request)
	}
}
