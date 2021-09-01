package example

import (
	"context"
	"os"

	"go.uber.org/zap"

	"github.com/zs368/gin-example/pkg/transport"
)

// Option is an application option.
type Option func(o *options)

// options is an application options.
type options struct {
	ctx  context.Context
	sigs []os.Signal

	logger  *zap.Logger
	servers []transport.Server
}

// Context with service context.
func Context(ctx context.Context) Option {
	return func(o *options) { o.ctx = ctx }
}

// Signal with exit signals.
func Signal(sigs ...os.Signal) Option {
	return func(o *options) { o.sigs = sigs }
}

// Logger with service logger.
func Logger(logger *zap.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

// Server with transport servers.
func Server(srv ...transport.Server) Option {
	return func(o *options) { o.servers = srv }
}
