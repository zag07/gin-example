package trace

import (
	"io"
	"time"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// FIXME 这儿是放在初始化好，还是中间件好？？？
func InitGlobalTracer(serviceName, agentHostPort string) (io.Closer, error) {
	cfg := &jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHostPort,
		},
	}

	closer, err := cfg.InitGlobalTracer(serviceName)
	if err != nil {
		return nil, err
	}
	// defer closer.Close()

	return closer, nil
}
