package service

import (
	pb "github.com/zs368/gin-example/api/example/v1"
	"github.com/zs368/gin-example/internal/biz"
	"go.uber.org/zap"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewBlogService)

type BlogService struct {
	pb.UnimplementedBlogServiceServer

	log *zap.Logger

	article *biz.ArticleUsecase
}

