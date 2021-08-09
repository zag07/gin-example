//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/zs368/gin-example"
	"github.com/zs368/gin-example/pkg/database"
	"github.com/zs368/gin-example/pkg/log"
	"github.com/zs368/gin-example/pkg/trace"
)

func initApp() (*gin_example.App, error) {
	panic(wire.Build(
		database.SetDB,
		log.CustomLogger,
		trace.InitGlobalTracer,
		gin_example.AppSet,
	))
}
