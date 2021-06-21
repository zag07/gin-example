package tools

import (
	"github.com/gin-gonic/gin"
	"github.com/zs368/gin-example/configs"
)

func GetPage(ctx *gin.Context) int {
	page := StrTo(ctx.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(ctx *gin.Context) int {
	pageSize := StrTo(ctx.Query("limit")).MustInt()
	if pageSize <= 0 {
		return configs.App.DefaultPageSize
	}
	if pageSize > configs.App.MaxPageSize {
		return configs.App.MaxPageSize
	}

	return pageSize
}

// TODO 下面这种形式是错误的 0.0
// func GetPageOffset(ctx *gin.Context, ...[]int) int {
//
// }
