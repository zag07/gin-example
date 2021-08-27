// 感觉这儿应该拆成article tag的，先不管了
package service

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/zs368/gin-example/internal/pkg/app"
	"github.com/zs368/gin-example/internal/pkg/errcode"
	"github.com/zs368/gin-example/internal/router/rule"
)

// 这儿的定义方式，要不要先预先定义，通过api层进行参数校验

func (s ExampleService) GetArticle(ctx *gin.Context) {
	r := app.NewResponse(ctx)
	params := rule.ArticleGetRequest{ID: cast.ToUint(ctx.Param("id"))}

	if err := app.BindAndValid(ctx, &params); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

	res, err := s.blog.GetArticle(ctx, cast.ToInt64(ctx.Param("id")))
	if err != nil {
		r.ToErrorResponse(errcode.ServerError)
		return
	}

	r.ToResponse(res)
}
