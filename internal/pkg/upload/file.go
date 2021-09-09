package upload

import (
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/zag07/gin-example/internal/conf"
	"github.com/zag07/gin-example/internal/pkg/utils"
)

type File struct {
	cfg *conf.HTTP
}

func NewFile(cfg *conf.HTTP) *File {
	return &File{cfg: cfg}
}

func (f *File) Cfg() *conf.HTTP {
	return f.cfg
}

type FileType int

const TypeImage FileType = iota + 1

func (f *File) GetFileName(name string) string {
	ext := path.Ext(name)
	fileName := utils.EncodeMD5(strings.TrimSuffix(name, ext))

	return fileName + ext
}

func (f *File) GetSavePath() string {
	if f.cfg.UploadSavePath != "" {
		return f.cfg.UploadSavePath
	}
	return "storage/app/uploads"
}

func (f *File) ISErrExist(dst string) bool {
	_, err := os.Stat(dst)
	return errors.Is(err, os.ErrExist)
}

func (f *File) IsErrPermission(dst string) bool {
	_, err := os.Stat(dst)
	return errors.Is(err, os.ErrPermission)
}

func (f *File) CheckContainExt(t FileType, name string) bool {
	ext := strings.ToUpper(path.Ext(name))
	switch t {
	case TypeImage:
		for _, allowExt := range f.cfg.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}

	return false
}

func (f *File) CheckFileSize(t FileType, file multipart.File) bool {
	content, _ := ioutil.ReadAll(file)
	size := int32(len(content))
	switch t {
	case TypeImage:
		if size <= f.cfg.UploadImageMaxSize<<20 {
			return true
		}
	}

	return false
}

func (f *File) CreateSavePath(path string, perm os.FileMode) error {
	if err := os.MkdirAll(path, perm); err != nil {
		return err
	}

	return nil
}

func (f *File) SaveFile(file *multipart.FileHeader, path string) error {
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
