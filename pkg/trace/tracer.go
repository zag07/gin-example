package trace

import (
	"io"
	"time"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func InitGlobalTracer() (io.Closer, error) {
	cfg := &jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "127.0.0.1:6831",
		},
	}

	closer, err := cfg.InitGlobalTracer("gin-example")
	if err != nil {
		return closer, err
	}
	// defer closer.Close()

	return closer, nil
}
