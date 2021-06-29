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
		ID uint32 `uri:"id" binding:"required,numeric"`
	}{}

	var (
		r   = core.NewResponse(c)
		err error
	)
	if err = c.ShouldBindUri(&params); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	var (
		db      = database.DB
		article models.Article
	)
	if err = db.Where("id = ?", params.ID).First(&article).Error; err != nil {
		r.ToErrorResponse(errcode.ArticleGetFail)
		return
	}

	r.ToResponse(article)
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

	var (
		r   = core.NewResponse(c)
		err error
	)
	if err = c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	var (
		db      = database.DB
		article = models.Article{
			Title:         params.Title,
			Desc:          params.Desc,
			CoverImageUrl: params.CoverImageUrl,
			Content:       params.Content,
			State:         params.State,
			CreatedBy:     params.CreatedBy,
		}
	)
	if err = db.Create(&article).Error; err != nil {
		r.ToErrorResponse(errcode.ArticleCreateFail)
		return
	}

	r.ToResponse(article)
}

func (a Article) Update(c *gin.Context) {
	params := struct {
		ID            uint   `json:"id" binding:"required,gte=1"`
		Title         string `json:"title" binding:"min=2,max=100"`
		Desc          string `json:"desc" binding:"min=2,max=255"`
		Content       string `json:"content" binding:"min=2,max=4294967295"`
		CoverImageUrl string `json:"cover_image_url" binding:"omitempty,url"`
		State         uint8  `json:"state" binding:"oneof=0 1"`
		UpdatedBy     string `json:"updated_by" binding:"required,min=2,max=100"`
	}{}

	var (
		r   = core.NewResponse(c)
		err error
	)
	if err = c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	db := database.DB
	if err = db.Model(models.Article{}).Where("id = ?", params.ID).Updates(models.Article{
		Title:         params.Title,
		Desc:          params.Desc,
		Content:       params.Content,
		CoverImageUrl: params.CoverImageUrl,
		State:         params.State,
		UpdatedBy:     params.UpdatedBy,
	}).Error; err != nil {
		r.ToErrorResponse(errcode.ArticleUpdateFail)
		return
	}

	r.ToResponse("文章更新成功")
}

func (a Article) Delete(c *gin.Context) {
	params := struct {
		ID uint32 `json:"id" binding:"required,gte=1"`
	}{}

	var (
		r   = core.NewResponse(c)
		err error
	)
	if err = c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	db := database.DB
	if err = db.Delete(&models.Article{}, params.ID).Error; err != nil {
		r.ToErrorResponse(errcode.ArticleDeleteFail)
		return
	}

	r.ToResponse("文章删除成功")
}
