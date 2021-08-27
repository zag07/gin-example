package biz

import (
	"context"
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
	log  *zap.Logger
}

func NewBlogUseCase(repo BlogRepo, logger *zap.Logger) *BlogUseCase {
	return &BlogUseCase{repo: repo}
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
