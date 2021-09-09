//go:build wireinject

package main

import (
	"github.com/google/wire"
	"go.uber.org/zap"

	"github.com/zag07/gin-example/internal/biz"
	"github.com/zag07/gin-example/internal/conf"
	"github.com/zag07/gin-example/internal/data"
	"github.com/zag07/gin-example/internal/server"
	"github.com/zag07/gin-example/internal/service"
)

// initApp init example application.
func initApp(*conf.HTTP, *conf.Data, *zap.Logger) (*example.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
