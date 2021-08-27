package data

import (
	"context"
	"time"

	"github.com/zs368/gin-example/internal/biz"
	"go.uber.org/zap"
)

type blogRepo struct {
	data *Data
	log  *zap.Logger
}

// NewBlogRepo .
func NewBlogRepo(data *Data, logger *zap.Logger) biz.BlogRepo {
	return &blogRepo{
		data: data,
		log:  logger,
	}
}

func (ar *blogRepo) ListArticle(ctx context.Context) ([]*biz.Article, error) {
	ps, err := ar.data.db.Article.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Article, 0)
	for _, p := range ps {
		rv = append(rv, &biz.Article{
			Id:        p.ID,
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	return rv, nil
}

func (ar *blogRepo) GetArticle(ctx context.Context, id int64) (*biz.Article, error) {
	p, err := ar.data.db.Article.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Article{
		Id:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, nil
}

func (ar *blogRepo) CreateArticle(ctx context.Context, article *biz.Article) error {
	_, err := ar.data.db.Article.
		Create().
		SetTitle(article.Title).
		SetContent(article.Content).
		Save(ctx)
	return err
}

func (ar *blogRepo) UpdateArticle(ctx context.Context, id int64, article *biz.Article) error {
	p, err := ar.data.db.Article.Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetTitle(article.Title).
		SetContent(article.Content).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	return err
}

func (ar *blogRepo) DeleteArticle(ctx context.Context, id int64) error {
	return ar.data.db.Article.DeleteOneID(id).Exec(ctx)
}
