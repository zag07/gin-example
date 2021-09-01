// 感觉这儿应该拆成article tag的，先不管了
package service

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/zs368/gin-example/internal/biz"
	"github.com/zs368/gin-example/internal/pkg/app"
	"github.com/zs368/gin-example/internal/pkg/errcode"
	"github.com/zs368/gin-example/internal/router/rule"
)

// 这儿的定义方式，要不要先预先定义，通过api层进行参数校验

func (s *ExampleService) GetArticle(ctx *gin.Context) {
	r := app.NewResponse(ctx)
	p := rule.GetArticle{ID: cast.ToInt64(ctx.Param("id"))}

	if err := app.BindAndValid(ctx, &p); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

	res, err := s.blog.GetArticle(ctx, p.ID)
	if err != nil {
		r.ToErrorResponse(errcode.GetArticle.WithDetails(err.Error()))
		return
	}

	r.ToResponse(res)
}

func (s *ExampleService) ListArticle(ctx *gin.Context) {
	r := app.NewResponse(ctx)
	p := rule.ListArticle{}

	if err := app.BindAndValid(ctx, &p); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

	res, err := s.blog.ListArticle(ctx)
	if err != nil {
		r.ToErrorResponse(errcode.ListArticle.WithDetails(err.Error()))
		return
	}

	r.ToResponse(res)
}

func (s *ExampleService) CreateArticle(ctx *gin.Context) {
	r := app.NewResponse(ctx)
	p := rule.CreateArticle{}

	if err := app.BindAndValid(ctx, &p); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

	err := s.blog.CreateArticle(ctx, &biz.Article{
		Title:         p.Title,
		Desc:          p.Desc,
		CoverImageUrl: p.CoverImageUrl,
		Content:       p.Content,
		CreatedBy:     p.CreatedBy,
		UpdatedBy:     p.UpdatedBy,
	})
	if err != nil {
		r.ToErrorResponse(errcode.CreateArticle.WithDetails(err.Error()))
		return
	}

	r.ToResponse("文章创建成功")
}

func (s *ExampleService) UpdateArticle(ctx *gin.Context) {
	r := app.NewResponse(ctx)
	p := rule.UpdateArticle{ID: cast.ToInt64(ctx.Param("id"))}

	if err := app.BindAndValid(ctx, &p); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

	err := s.blog.UpdateArticle(ctx, p.ID, &biz.UpdateArticle{
		Title:         p.Title,
		Desc:          p.Desc,
		CoverImageUrl: p.CoverImageUrl,
		Content:       p.Content,
		Status:        p.Status,
		UpdatedBy:     p.UpdatedBy,
	})
	if err != nil {
		r.ToErrorResponse(errcode.Err.WithDetails(err.Error()))
		return
	}
	r.ToResponse("文章更新成功")
}

func (s *ExampleService) DeleteArticle(ctx *gin.Context) {
	r := app.NewResponse(ctx)
	p := rule.DeleteArticle{ID: cast.ToInt64(ctx.Param("id"))}

	if err := app.BindAndValid(ctx, &p); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

	err := s.blog.DeleteArticle(ctx, p.ID)
	if err != nil {
		r.ToErrorResponse(errcode.DeleteArticle.WithDetails(err.Error()))
	}

	r.ToResponse("文章删除成功")
}

func (s *ExampleService) UploadFile(ctx *gin.Context) {
	r := app.NewResponse(ctx)
	p := rule.UploadFile{Type: cast.ToUint(ctx.Param("type"))}

	if err := app.BindAndValid(ctx, &p); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

}

// TODO tag 类似 article
