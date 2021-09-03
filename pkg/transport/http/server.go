package http

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/zs368/gin-example/pkg/middleware"
)

// Server is an HTTP server wrapper.
type Server struct {
	*http.Server
	ms      []middleware.Middleware
	timeout time.Duration
	log     *zap.Logger
}

// ServerOption is an HTTP server option.
type ServerOption func(*Server)

// Addr with server address.
func Addr(addr string) ServerOption {
	return func(s *Server) {
		s.Addr = addr
	}
}

// Handler with server handler.
func Handler(handler http.Handler) ServerOption {
	return func(s *Server) {
		s.Handler = handler
	}
}

// Timeout with server timeout.
func Timeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

// Middleware with service middleware option.
func Middleware(m ...middleware.Middleware) ServerOption {
	return func(s *Server) {
		s.ms = m
	}
}

// Logger with server logger.
func Logger(logger *zap.Logger) ServerOption {
	return func(s *Server) {
		s.log = logger
	}
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		Server: &http.Server{
			Addr: ":0",
		},
		timeout: 1 * time.Second,
		log:     zap.NewExample(),
	}
	for _, opt := range opts {
		opt(srv)
	}
	return srv
}

func (s *Server) Start(ctx context.Context) error {
	s.log.Sugar().Info("http: starting web server at %s", s.Addr)

	s.BaseContext = func(net.Listener) context.Context {
		return ctx
	}

	err := s.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.Shutdown(ctx)
}
