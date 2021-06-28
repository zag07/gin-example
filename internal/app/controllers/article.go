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
		Title         string `json:"title" binding:"required,min=2,max=100"`
		Desc          string `json:"desc" binding:"required,min=2,max=255"`
		Content       string `json:"content" binding:"required,min=2,max=4294967295"`
		CoverImageUrl string `json:"cover_image_url" binding:"omitempty,url"`
		CreatedBy     string `json:"created_by" binding:"required,min=2,max=100"`
		State         uint8  `json:"state" binding:"oneof=0 1"`
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
	params := struct {
		ID            uint32 `json:"id" binding:"required,gte=1"`
		TagID         uint32 `json:"tag_id" binding:"required,gte=1"`
		Title         string `json:"title" binding:"min=2,max=100"`
		Desc          string `json:"desc" binding:"min=2,max=255"`
		Content       string `json:"content" binding:"min=2,max=4294967295"`
		CoverImageUrl string `json:"cover_image_url" binding:"omitempty,url"`
		ModifiedBy    string `json:"modified_by" binding:"required,min=2,max=100"`
		State         uint8  `json:"state" binding:"oneof=0 1"`
	}{}

	r := core.NewResponse(c)
	var err error

	if err = c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	err = models.Article{ID: params.ID}.Update(database.DB, map[string]interface{}{
		"state": params.State,
	})

	if err != nil {
		r.ToErrorResponse(errcode.ArticleUpdateFail)
	}

	r.ToResponse("更新成功")
}

func (a Article) Delete(c *gin.Context) {

}
