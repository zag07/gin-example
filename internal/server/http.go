package server

import (
	"go.uber.org/zap"

	"github.com/zs368/gin-example/internal/conf"
	"github.com/zs368/gin-example/internal/router"
	"github.com/zs368/gin-example/internal/service"
	"github.com/zs368/gin-example/pkg/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(cfg *conf.HTTP, logger *zap.Logger, e *service.ExampleService) *http.Server {
	var opts = []http.ServerOption{
		http.Addr(cfg.Port),
		http.Handler(router.NewRouter(cfg)),
		http.Logger(logger),
	}
	srv := http.NewServer(opts...)
	return srv
}
