package service

import (
	"github.com/zag07/gin-example/internal/pkg/chatroom"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zag07/gin-example/internal/pkg/app"
	"github.com/zag07/gin-example/internal/pkg/errcode"
)

var broadcast = chat_svs.Broadcaster

func (s *ExampleService) Home(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request

	http.ServeFile(w, r, "./assets/template/home.html")
}

func (s *ExampleService) Home2(ctx *gin.Context) {
	r := app.NewResponse(ctx)
	w := ctx.Writer

	tpl, err := template.ParseFiles("./assets/template/home2.html")
	if err != nil {
		r.ToErrorResponse(errcode.Err.WithDetails("模板解析错误！"))
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		r.ToErrorResponse(errcode.Err.WithDetails("模板执行错误！"))
		return
	}
}

func (s *ExampleService) UserList(ctx *gin.Context) {
	r := app.NewResponse(ctx)
	var userList = broadcast.GetUserList()

	r.ToResponse(userList)
}

func (s *ExampleService) WS(ctx *gin.Context) {
	chat_svs.ServeWs(ctx.Writer, ctx.Request)
}
