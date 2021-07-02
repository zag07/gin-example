package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zs368/gin-example/internal/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r Response) ToResponseList(list interface{}, total int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"paper": Pager{
			Page:     GetPage(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			Total:    total,
		},
	})
}

func (r Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"errcode": err.Code(), "errmsg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}
