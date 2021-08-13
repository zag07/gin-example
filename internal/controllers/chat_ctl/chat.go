package chat_ctl

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zs368/gin-example/internal/pkg/app"
	"github.com/zs368/gin-example/internal/pkg/errcode"
	"github.com/zs368/gin-example/internal/services/chat_svs"
)

type Chat struct{}

func NewChat() Chat {
	return Chat{}
}

var broadcast = chat_svs.Broadcaster

func (c Chat) Home(ctx *gin.Context) {
	var (
		w = ctx.Writer
		r = ctx.Request
	)

	http.ServeFile(w, r, "./assets/template/home.html")
}

func (c Chat) Home2(ctx *gin.Context) {
	var (
		r = app.NewResponse(ctx)
		w = ctx.Writer
	)

	tpl, err := template.ParseFiles("./assets/template/home2.html")
	if err != nil {
		r.ToErrorResponse(errcode.ServerError.WithDetails("模板解析错误！"))
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		r.ToErrorResponse(errcode.ServerError.WithDetails("模板执行错误！"))
		return
	}
}

func (c Chat) UserList(ctx *gin.Context) {
	var (
		r        = app.NewResponse(ctx)
		userList = broadcast.GetUserList()
	)

	r.ToResponse(userList)
}

func (c Chat) WS(ctx *gin.Context) {
	chat_svs.ServeWs(ctx.Writer, ctx.Request)
}
