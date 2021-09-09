package data

import (
	"context"
	"github.com/zag07/gin-example/internal/biz"
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
		SetDesc(article.Desc).
		SetCoverImageURL(article.CoverImageUrl).
		SetNillableContent(&article.Content).
		SetCreatedBy(article.CreatedBy).
		SetUpdatedBy(article.UpdatedBy).
		Save(ctx)
	return err
}

func (ar *blogRepo) UpdateArticle(ctx context.Context, id int64, article *biz.UpdateArticle) error {
	p, err := ar.data.db.Article.Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetNillableTitle(article.Title).
		SetNillableDesc(article.Desc).
		SetNillableCoverImageURL(article.CoverImageUrl).
		SetNillableContent(article.Content).
		SetNillableStatus(article.Status).
		SetNillableUpdatedBy(article.UpdatedBy).
		Save(ctx)
	return err
}

func (ar *blogRepo) DeleteArticle(ctx context.Context, id int64) error {
	return ar.data.db.Article.DeleteOneID(id).Exec(ctx)
}
