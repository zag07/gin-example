package server

import (
	"go.uber.org/zap"

	"github.com/zag07/gin-example/internal/conf"
	"github.com/zag07/gin-example/internal/router"
	"github.com/zag07/gin-example/internal/service"
	"github.com/zag07/gin-example/pkg/middleware/tracing"
	"github.com/zag07/gin-example/pkg/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(cfg *conf.HTTP, logger *zap.Logger, e *service.ExampleService) *http.Server {
	var opts = []http.ServerOption{
		http.Addr(cfg.Port),
		http.Handler(router.NewRouter(cfg, e)),
		http.Logger(logger),
		http.Middleware(
			tracing.Server(),
		),
	}
	srv := http.NewServer(opts...)
	return srv
}
