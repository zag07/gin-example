package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/zs368/gin-example/internal/app/models"
	"github.com/zs368/gin-example/internal/app/rules"
	"github.com/zs368/gin-example/internal/pkg/app"
	"github.com/zs368/gin-example/internal/pkg/errcode"
	"github.com/zs368/gin-example/pkg/database"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// Get godoc
// @Summary 获取单个文章
// @Produce  json
// @Param id path int true "文章 ID"
// @Success 200 {array} rules.ArticleGetResponse "获取文章成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article/{id} [get]
func (a Article) Get(c *gin.Context) {
	var (
		r      = app.NewResponse(c)
		params = rules.ArticleGetRequest{}
	)

	if err := app.BindAndValid(c, &params); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

	var (
		db      = database.DB
		article models.Article
		res     rules.ArticleGetResponse
	)

	if err := db.Model(&article).Where("id = ?", params.ID).First(&res).Error; err != nil {
		r.ToErrorResponse(errcode.ArticleGetFail.WithDetails(err.Error()))
		return
	}

	r.ToResponse(res)
}

// Create godoc
// @Summary 创建文章
// @Accept  json
// @Produce json
// @Param title body string true "文章标题"
// @Param desc body string true "文章简述"
// @Param cover_image_url body string false "封面图片地址"
// @Param content body string true "文章内容"
// @Param created_by body int true "创建者"
// @Param updated_by body int true "创建者"
// @Success 200 {string} string "文章创建成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article [post]
func (a Article) Create(c *gin.Context) {
	var (
		r      = app.NewResponse(c)
		params = rules.ArticleCreateRequest{}
	)

	if err := app.BindAndValid(c, &params); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

	var (
		db      = database.DB
		article = models.Article{
			Title:         params.Title,
			Desc:          params.Desc,
			CoverImageUrl: params.CoverImageUrl,
			Content:       params.Content,
			CreatedBy:     params.CreatedBy,
			UpdatedBy:     params.UpdatedBy,
		}
	)
	if err := db.Create(&article).Error; err != nil {
		r.ToErrorResponse(errcode.ArticleCreateFail.WithDetails(err.Error()))
		return
	}

	r.ToResponse("文章创建成功")
}

// Update godoc
// @Summary 更新文章
// @Accept json
// @Produce json
// @Param id path int true "文章ID"
// @Param title body string false "文章标题"
// @Param desc body string false "文章简述"
// @Param cover_image_url body string false "封面图片地址"
// @Param content body string false "文章内容"
// @Param modified_by body string true "修改者"
// @Success 200 {string} string "文章更新成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article/{id} [put]
func (a Article) Update(c *gin.Context) {
	var (
		r      = app.NewResponse(c)
		params = rules.ArticleUpdateRequest{ID: cast.ToUint(c.Param("id"))}
	)

	if err := app.BindAndValid(c, &params); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

	var (
		db      = database.DB
		article models.Article
		data    = map[string]interface{}{}
	)

	if params.Title != "" {
		data["title"] = params.Title
	}
	if params.Desc != "" {
		data["desc"] = params.Desc
	}
	if params.Content != "" {
		data["content"] = params.Content
	}
	if params.CoverImageUrl != "" {
		data["cover_image_url"] = params.CoverImageUrl
	}
	// TODO 状态为0的时候
	if params.State != 0 {
		data["state"] = params.State
	}
	if params.CoverImageUrl != "" {
		data["updated_by"] = params.CoverImageUrl
	}
	// TODO id 未查到时，没有数据更新时
	if err := db.Model(&article).Where("id = ?", params.ID).Updates(data).Error; err != nil {
		r.ToErrorResponse(errcode.ArticleUpdateFail.WithDetails(err.Error()))
		return
	}

	r.ToResponse("文章更新成功")
}

// Delete godoc
// @Summary 删除文章
// @Produce  json
// @Param id path int true "文章ID"
// @Success 200 {string} string "文章删除成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	var (
		r      = app.NewResponse(c)
		params = rules.ArticleDeleteRequest{}
	)

	if err := app.BindAndValid(c, &params); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

	var (
		db      = database.DB
		article models.Article
	)

	if err := db.Delete(&article, params.ID).Error; err != nil {
		r.ToErrorResponse(errcode.ArticleDeleteFail.WithDetails(err.Error()))
		return
	}

	r.ToResponse("文章删除成功")
}
