package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/cast"
	"github.com/zs368/gin-example/internal/app/controllers/core"
	"github.com/zs368/gin-example/internal/app/models"
	"github.com/zs368/gin-example/internal/pkg/errcode"
	"github.com/zs368/gin-example/pkg/database"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// Get @Summary 获取多个标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Success 200 {object} models.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tag/{id} [get]
func (t Tag) Get(c *gin.Context) {
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
		db  = database.DB
		tag models.Tag
	)
	if err = db.Where("id = ?", params.ID).First(&tag).Error; err != nil {
		r.ToErrorResponse(errcode.TagGetFail.WithDetails(err.Error()))
		return
	}

	r.ToResponse(tag)
}

// Create @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(2) maxlength(100)
// @Param state body int false "状态" Enums(0, 1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} models.Tag
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tag [post]
func (t Tag) Create(c *gin.Context) {
	params := struct {
		Name      string `json:"title" binding:"required,min=2,max=100"`
		State     uint8  `json:"state" binding:"oneof=0 1"`
		CreatedBy string `json:"created_by" binding:"required,min=2,max=100"`
		UpdatedBy string `json:"updated_by" binding:"required,min=2,max=100"`
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
		db  = database.DB
		tag = models.Tag{
			Name:      params.Name,
			State:     params.State,
			CreatedBy: params.CreatedBy,
			UpdatedBy: params.UpdatedBy,
		}
	)
	if err = db.Create(&tag).Error; err != nil {
		r.ToErrorResponse(errcode.TagCreateFail.WithDetails(err.Error()))
		return
	}

	r.ToResponse(tag)
}

// Update @Summary 更新标签
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
	params := struct {
		ID        uint   `json:"id" binding:"required,gte=1"`
		Name      string `json:"title" binding:"required,min=2,max=100"`
		State     uint8  `json:"state" binding:"oneof=0 1"`
		CreatedBy string `json:"created_by" binding:"required,min=2,max=100"`
		UpdatedBy string `json:"updated_by" binding:"required,min=2,max=100"`
	}{
		ID: cast.ToUint(c.Param("id")),
	}

	var (
		r   = core.NewResponse(c)
		err error
	)
	if err = c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		r.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	db := database.DB
	if err = db.Model(models.Tag{}).Where("id = ?", params.ID).Updates(models.Tag{
		Name:      params.Name,
		State:     params.State,
		CreatedBy: params.CreatedBy,
		UpdatedBy: params.UpdatedBy,
	}).Error; err != nil {
		r.ToErrorResponse(errcode.TagUpdateFail)
		return
	}

	r.ToResponse("标签更新成功")
}

// Delete @Summary 删除标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Success 200 {string} string "标签删除成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tag/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
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
	if err = db.Delete(&models.Tag{}, params.ID).Error; err != nil {
		r.ToErrorResponse(errcode.TagDeleteFail.WithDetails(err.Error()))
		return
	}

	r.ToResponse("标签删除成功")
}
