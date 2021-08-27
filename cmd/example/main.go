package main

import (
	"flag"
	"github.com/zs368/gin-example/pkg/transport/http"
	"go.uber.org/zap"

	example "github.com/zs368/gin-example"
)

var (
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "./configs", "config path, eg: -conf config.yaml")
}

func newApp(logger *zap.Logger, hs *http.Server) *example.App {
	return example.New(
		example.Logger(logger),
		example.Server(hs),
	)
}

// @title gin-example
// @version 0.2.x
/*func main() {
	flag.Parse()
	logger, err := log.CustomLogger()
	if err != nil {
		panic(err)
	}

	cfg, err := config.Load(flagconf)
	if err != nil {
		panic(err)
	}

	app, cleanup, err := initApp()

}
*/