package common_rule

// TODO
type UploadFileRequest struct {
	Type uint   `form:"type" binding:"required"`
	File string `file:"file" binding:"required"`
}

type UploadFileResponse struct {
	FileUrl string `json:"file_url"`
}
