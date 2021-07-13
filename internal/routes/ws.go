package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zs368/gin-example/internal/app/controllers/chat_ctl"
	"github.com/zs368/gin-example/internal/app/services/chat_svs"
)

func SetWSRouter(r *gin.Engine) {
	broadcaster := chat_svs.NewBroadcaster()
	go broadcaster.Run()

	c := r.Group("/chat")
	{
		chat := chat_ctl.NewChat()
		c.GET("/", chat.Home)
		c.GET("/ws", chat.WS(broadcaster))
	}
}
