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

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// Get godoc
// @Summary 获取单个标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Success 200 {array} rules.TagGetResponse "获取标签成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tag/{id} [get]
func (t Tag) Get(c *gin.Context) {
	var (
		r      = app.NewResponse(c)
		params = rules.TagGetRequest{}
	)

	if err := app.BindAndValid(c, &params); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

	var (
		db  = database.DB
		tag models.Tag
		res rules.TagGetResponse
	)

	if err := db.Model(&tag).Where("id = ?", params.ID).First(&res).Error; err != nil {
		r.ToErrorResponse(errcode.TagGetFail.WithDetails(err.Error()))
		return
	}

	r.ToResponse(res)
}

// Create godoc
// @Summary 新增标签
// @Accept  json
// @Produce  json
// @Param name body string true "标签名称" minlength(2) maxlength(100)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Param updated_by body string true "更新者" minlength(3) maxlength(100)
// @Success 200 {string} string "标签创建成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tag [post]
func (t Tag) Create(c *gin.Context) {
	var (
		r      = app.NewResponse(c)
		params = rules.TagCreateRequest{}
	)

	if err := app.BindAndValid(c, &params); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

	var (
		db  = database.DB
		tag = models.Tag{
			Name:      params.Name,
			CreatedBy: params.CreatedBy,
			UpdatedBy: params.UpdatedBy,
		}
	)

	if err := db.Create(&tag).Error; err != nil {
		r.ToErrorResponse(errcode.TagCreateFail.WithDetails(err.Error()))
		return
	}

	r.ToResponse("标签创建成功")
}

// Update godoc
// @Summary 更新标签
// @Accept  json
// @Produce  json
// @Param id path int true "标签 ID"
// @Param name body string false "标签名称" minlength(2) maxlength(100)
// @Param state body int false "状态" Enums(0, 1)
// @Param modified_by body string true "修改者" minlength(2) maxlength(100)
// @Success 200 {string} string "标签更新成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tag/{id} [put]
func (t Tag) Update(c *gin.Context) {
	var (
		r      = app.NewResponse(c)
		params = rules.TagUpdateRequest{ID: cast.ToUint(c.Param("id"))}
	)

	if err := app.BindAndValid(c, &params); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

	var (
		db  = database.DB
		tag models.Tag
	)

	if err := db.Model(&tag).Where("id = ?", params.ID).Updates(models.Tag{
		Name:      params.Name,
		CreatedBy: params.CreatedBy,
		UpdatedBy: params.UpdatedBy,
	}).Error; err != nil {
		r.ToErrorResponse(errcode.TagUpdateFail)
		return
	}

	r.ToResponse("标签更新成功")
}

// Delete
// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Success 200 {string} string "标签删除成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tag/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	var (
		r      = app.NewResponse(c)
		params = rules.TagDeleteRequest{}
	)

	if err := app.BindAndValid(c, &params); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

	var (
		db  = database.DB
		tag models.Tag
	)

	if err := db.Delete(&tag, params.ID).Error; err != nil {
		r.ToErrorResponse(errcode.TagDeleteFail.WithDetails(err.Error()))
		return
	}

	r.ToResponse("标签删除成功")
}
