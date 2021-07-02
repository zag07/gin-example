package app

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/zs368/gin-example/configs"
)

func GetPage(ctx *gin.Context) int {
	page := cast.ToInt(ctx.Query("page"))
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(ctx *gin.Context) int {
	pageSize := cast.ToInt(ctx.Query("limit"))
	if pageSize <= 0 {
		return configs.App.DefaultPageSize
	}
	if pageSize > configs.App.MaxPageSize {
		return configs.App.MaxPageSize
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	pageOffset := 0
	if page > 0 {
		pageOffset = (page - 1) * pageSize
	}

	return pageOffset
}
