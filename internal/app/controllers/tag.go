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
