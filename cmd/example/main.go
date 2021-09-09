package main

import (
	"flag"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.uber.org/zap"

	example "github.com/zag07/gin-example"
	"github.com/zag07/gin-example/pkg/config"
	"github.com/zag07/gin-example/pkg/log"
	"github.com/zag07/gin-example/pkg/transport/http"
)

var (
	// Name is the name of the compiled software.
	Name = "gin.example"
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "./configs/config.yaml", "config path, eg: -conf config.yaml")
}

func newApp(logger *zap.Logger, hs *http.Server) *example.App {
	return example.New(
		example.Logger(logger),
		example.Server(hs),
	)
}

func setTracerProvider(url string) error {
	// Set global trace provider
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return err
	}
	tp := tracesdk.NewTracerProvider(
		// Set the sampling rate based on the parent span to 100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(Name),
			attribute.String("env", "dev"),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}

// @title gin-example
// @version 0.2.x
func main() {
	flag.Parse()
	logger, err := log.CustomLogger()
	if err != nil {
		panic(err)
	}

	cfg, err := config.Load(flagconf)
	if err != nil {
		panic(err)
	}

	if err := setTracerProvider(cfg.Trace.Endpoint); err != nil {
		panic(err)
	}

	app, cleanup, err := initApp(cfg.Http, cfg.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
