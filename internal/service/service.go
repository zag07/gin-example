package service

import (
	"github.com/google/wire"
	"go.uber.org/zap"

	"github.com/zag07/gin-example/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewBlogService)

type ExampleService struct {
	blog *biz.BlogUseCase
	log  *zap.Logger
}

func NewBlogService(bc *biz.BlogUseCase, logger *zap.Logger) *ExampleService {
	return &ExampleService{
		blog: bc,
		log:  logger,
	}
}
