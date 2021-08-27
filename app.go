package example

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type App struct {
	ctx    context.Context
	opts   options
	cancel func()
}

func New(opts ...Option) *App {
	options := options{
		ctx:    context.Background(),
		logger: zap.NewExample(),
		sigs:   []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}

	ctx, cancel := context.WithCancel(options.ctx)
	return &App{
		ctx:    ctx,
		opts:   options,
		cancel: cancel,
	}
}

// Run executes all OnStart hooks registered with the application's Lifecycle.
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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, a.opts.sigs...)
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-quit:
				err := a.Stop()
				if err != nil {
					a.opts.logger.Sugar().Errorf("failed to app stop: %v", err)
				}
			}
		}
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

// Stop gracefully stops the application.
func (a *App) Stop() error {
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}
