package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/zs368/gin-example/pkg/tracer"
)

func Tracing() func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			newCtx context.Context
			span   opentracing.Span
		)

		spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		)
		if err != nil {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				tracer.Tracer,
				c.Request.URL.Path,
			)
		} else {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				tracer.Tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
			)
		}
		defer span.Finish()

		var (
			spanContext = span.Context()
			traceID     string
			spanID      string
		)

		switch spanContext.(type) {
		case jaeger.SpanContext:
			jaegerContext := spanContext.(jaeger.SpanContext)
			traceID = jaegerContext.TraceID().String()
			spanID = jaegerContext.SpanID().String()
		}
		c.Set("X-Trace-ID", traceID)
		c.Set("X-Span-ID", spanID)
		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}
