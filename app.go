package example

import (
	"context"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go.uber.org/zap"
)

type App struct {
	ctx    context.Context
	opts   options
	cancel func()
}

func New(opts ...Option) *App {
	options := options{
		ctx:    context.Background(),
		logger: zap.L(),
		sigs:   []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}

	ctx, cancel := context.WithCancel(options.ctx)
	return &App{
		ctx:    ctx,
		opts:   options,
		cancel: cancel,
	}
}

func (a *App) Run() error {
	eg, ctx := errgroup.WithContext(a.ctx)
	wg := sync.WaitGroup{}
	for _, srv := range a.opts.servers {
		srv := srv
		eg.Go(func() error {
			<-ctx.Done() // wait for stop signal
			return srv.Stop(ctx)
		})
		wg.Add(1)
		eg.Go(func() error {
			wg.Done()
			return srv.Start(ctx)
		})
	}
	wg.Wait()

	c := make(chan os.Signal, 1)
	signal.Notify(c, a.opts.sigs...)

}

func (a *App) Stop() error {

}
