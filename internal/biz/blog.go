package biz

import (
	"context"
	"errors"
	"github.com/zs368/gin-example/internal/conf"
	"github.com/zs368/gin-example/internal/pkg/upload"
	"mime/multipart"
	"os"
	"time"

	"go.uber.org/zap"
)

type Article struct {
	Id        int64
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Like      int64
}

type BlogRepo interface {
	// db
	ListArticle(ctx context.Context) ([]*Article, error)
	GetArticle(ctx context.Context, id int64) (*Article, error)
	CreateArticle(ctx context.Context, article *Article) error
	UpdateArticle(ctx context.Context, id int64, article *Article) error
	DeleteArticle(ctx context.Context, id int64) error

	// redis
	GetArticleLike(ctx context.Context, id int64) (rv int64, err error)
	IncArticleLike(ctx context.Context, id int64) error
}

type BlogUseCase struct {
	repo BlogRepo
	cfg  *conf.HTTP
	log  *zap.Logger
}

func NewBlogUseCase(repo BlogRepo, cfg *conf.HTTP, logger *zap.Logger) *BlogUseCase {
	return &BlogUseCase{repo: repo, cfg: cfg, log: logger}
}

func (uc *BlogUseCase) GetArticle(ctx context.Context, id int64) (p *Article, err error) {
	p, err = uc.repo.GetArticle(ctx, id)
	if err != nil {
		return
	}
	err = uc.repo.IncArticleLike(ctx, id)
	if err != nil {
		return
	}
	p.Like, err = uc.repo.GetArticleLike(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *BlogUseCase) ListArticle(ctx context.Context) (ps []*Article, err error) {
	ps, err = uc.repo.ListArticle(ctx)
	if err != nil {
		return
	}
	return
}

func (uc *BlogUseCase) CreateArticle(ctx context.Context, article *Article) error {
	return uc.repo.CreateArticle(ctx, article)
}

func (uc *BlogUseCase) UpdateArticle(ctx context.Context, id int64, article *Article) error {
	return uc.repo.UpdateArticle(ctx, id, article)
}

func (uc *BlogUseCase) DeleteArticle(ctx context.Context, id int64) error {
	return uc.repo.DeleteArticle(ctx, id)
}

type FileInfo struct {
	Name string
	Url  string
}

func (uc *BlogUseCase) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	f := upload.NewFile(uc.cfg)
	fileName := f.GetFileName(fileHeader.Filename)
	if !f.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if !f.CheckFileSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}

	uploadSavePath := f.GetSavePath()
	if f.ISErrExist(uploadSavePath) {
		if err := f.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if f.IsErrPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}
	if err := f.SaveFile(fileHeader, uploadSavePath+"/"+fileName); err != nil {
		return nil, err
	}

	return &FileInfo{
		Name: fileName,
		Url:  f.Cfg().UploadServerUrl + "/" + fileName,
	}, nil
}
