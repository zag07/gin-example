package common_ctl

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/zs368/gin-example/internal/pkg/app"
	"github.com/zs368/gin-example/internal/pkg/errcode"
	"github.com/zs368/gin-example/internal/pkg/upload"
	"github.com/zs368/gin-example/internal/rule/common_rule"
	"github.com/zs368/gin-example/internal/services/common_svs"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

// UploadFile godoc
// @Summary 文件上传
// @Produce  json
// @Param type body int true "文件类型"
// @Param file body string true "文件"
// @Success 200 {array} common_rule.UploadFileResponse "文件上传成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /c/upload/file [post]
func (u Upload) UploadFile(c *gin.Context) {
	var (
		r                     = app.NewResponse(c)
		fileType              = cast.ToInt(c.PostForm("type"))
		file, fileHeader, err = c.Request.FormFile("file")
	)

	if err != nil || fileHeader == nil || fileType <= 0 {
		r.ToErrorResponse(errcode.InvalidParams)
		return
	}

	fileInfo, err := common_svs.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		r.ToErrorResponse(errcode.UploadFileFail.WithDetails(err.Error()))
		return
	}

	r.ToResponse(common_rule.UploadFileResponse{FileUrl: fileInfo.Url})
}
