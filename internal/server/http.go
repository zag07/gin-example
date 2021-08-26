package server

import (
	"context"
	"net"
	"net/http"

	"go.uber.org/zap"

	"github.com/zs368/gin-example/internal/conf"
	"github.com/zs368/gin-example/internal/router"
)

type HTTPServer struct {
	*http.Server
	cfg *conf.HTTP
	log *zap.Logger
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(cfg *conf.HTTP, logger *zap.Logger) *HTTPServer {
	return &HTTPServer{
		Server: &http.Server{
			Addr:    cfg.Port,
			Handler: router.NewRouter(cfg),
		},
		cfg: cfg,
		log: logger,
	}
}

func (s *HTTPServer) Start(ctx context.Context) error {
	s.log.Sugar().Info("http: starting web server at %s", s.cfg.Port)

	s.BaseContext = func(net.Listener) context.Context {
		return ctx
	}
	return s.ListenAndServe()
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	return s.Shutdown(ctx)
}
