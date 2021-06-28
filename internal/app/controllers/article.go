package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/zs368/gin-example/internal/app/controllers/core"
	"github.com/zs368/gin-example/internal/app/models"
	"github.com/zs368/gin-example/internal/pkg/errcode"
	"github.com/zs368/gin-example/pkg/database"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	params := struct {
		ID string `uri:"id" binding:"required,numeric"`
	}{}

	r := core.NewResponse(c)
	if err := c.ShouldBindUri(&params); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	r.ToResponse("article get ok!!!")
}

func (a Article) Create(c *gin.Context) {
	params := struct {
		Title         string `form:"title" binding:"required,min=2,max=100"`
		Desc          string `form:"desc" binding:"required,min=2,max=255"`
		Content       string `form:"content" binding:"required,min=2,max=4294967295"`
		CoverImageUrl string `form:"cover_image_url" binding:"required,url"`
		CreatedBy     string `form:"created_by" binding:"required,min=2,max=100"`
		State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	}{}

	r := core.NewResponse(c)
	var err error

	if err = c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	article, err := models.Article{
		Title:         params.Title,
		Desc:          params.Desc,
		CoverImageUrl: params.CoverImageUrl,
		Content:       params.Content,
		State:         params.State,
		CreatedBy:     params.CreatedBy,
	}.Create(database.DB)

	r.ToResponse(article)
}

func (a Article) Update(c *gin.Context) {

}

func (a Article) Delete(c *gin.Context) {

}
