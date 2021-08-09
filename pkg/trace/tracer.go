package trace

import (
	"io"
	"time"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/zs368/gin-example/configs"
)

// FIXME 这儿是放在初始化好，还是中间件好？？？
func InitGlobalTracer() (io.Closer, error) {
	cfg := &jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  configs.App.TracePort,
		},
	}

	closer, err := cfg.InitGlobalTracer(configs.App.TraceName)
	if err != nil {
		return closer, err
	}
	// defer closer.Close()

	return closer, nil
}
