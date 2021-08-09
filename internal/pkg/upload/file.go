package upload

import (
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/zs368/gin-example/configs"
	"github.com/zs368/gin-example/internal/pkg/utils"
)

type FileType int

const TypeImage FileType = iota + 1

func GetFileName(name string) string {
	ext := path.Ext(name)
	fileName := utils.EncodeMD5(strings.TrimSuffix(name, ext))

	return fileName + ext
}

func GetSavePath() string {
	return configs.App.UploadSavePath
}

func ISErrExist(dst string) bool {
	_, err := os.Stat(dst)
	return errors.Is(err, os.ErrExist)
}

func IsErrPermission(dst string) bool {
	_, err := os.Stat(dst)
	return errors.Is(err, os.ErrPermission)
}

func CheckContainExt(t FileType, name string) bool {
	ext := strings.ToUpper(path.Ext(name))
	switch t {
	case TypeImage:
		for _, allowExt := range configs.App.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}

	return false
}

func CheckFileSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := int32(len(content))
	switch t {
	case TypeImage:
		if size <= configs.App.UploadImageMaxSize<<20 {
			return true
		}
	}

	return false
}

func CreateSavePath(path string, perm os.FileMode) error {
	if err := os.MkdirAll(path, perm); err != nil {
		return err
	}

	return nil
}

func SaveFile(file *multipart.FileHeader, path string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}
