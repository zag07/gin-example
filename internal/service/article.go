package service

import (
	"context"
	"fmt"

	pb "github.com/zs368/gin-example/api/example/v1"
	"github.com/zs368/gin-example/internal/biz"

	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

func NewBlogService(article *biz.ArticleUsecase, logger zap.Logger) *BlogService {
	return &BlogService{
		article: article,
		log:     &logger,
	}
}

func (s *BlogService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
	s.log.Info(fmt.Sprintf("input data %v", req))
	err := s.article.Create(ctx, &biz.Article{
		Title:   req.Title,
		Content: req.Content,
	})
	return &pb.CreateArticleReply{}, err
}

func (s *BlogService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	s.log.Info(fmt.Sprintf("input data %v", req))
	err := s.article.Update(ctx, req.Id, &biz.Article{
		Title:   req.Title,
		Content: req.Content,
	})
	return &pb.UpdateArticleReply{}, err
}

func (s *BlogService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
	s.log.Info(fmt.Sprintf("input data %v", req))
	err := s.article.Delete(ctx, req.Id)
	return &pb.DeleteArticleReply{}, err
}

func (s *BlogService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetArticle")
	defer span.End()
	p, err := s.article.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetArticleReply{Article: &pb.Article{Id: p.Id, Title: p.Title, Content: p.Content, Like: p.Like}}, nil
}

func (s *BlogService) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	ps, err := s.article.List(ctx)
	reply := &pb.ListArticleReply{}
	for _, p := range ps {
		reply.Results = append(reply.Results, &pb.Article{
			Id:      p.Id,
			Title:   p.Title,
			Content: p.Content,
		})
	}
	return reply, err
}
