package common_svs

import (
	"errors"
	"github.com/zs368/gin-example/configs"
	"mime/multipart"
	"os"

	"github.com/zs368/gin-example/internal/pkg/upload"
)

type FileInfo struct {
	Name string
	Url  string
}

func UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if !upload.CheckFileSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}

	uploadSavePath := upload.GetSavePath()
	if upload.ISErrExist(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.IsErrPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}
	if err := upload.SaveFile(fileHeader, uploadSavePath+"/"+fileName); err != nil {
		return nil, err
	}

	return &FileInfo{
		Name: fileName,
		Url:  configs.App.UploadServerUrl + "/" + fileName,
	}, nil
}
