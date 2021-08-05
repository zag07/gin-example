package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"gorm.io/gorm"
)

func Tracing(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			newCtx context.Context
			span   opentracing.Span
		)

		spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(ctx.Request.Header),
		)
		if err != nil {
			span, newCtx = opentracing.StartSpanFromContext(
				ctx.Request.Context(),
				ctx.Request.URL.Path,
			)
		} else {
			span, newCtx = opentracing.StartSpanFromContext(
				ctx.Request.Context(),
				ctx.Request.URL.Path,
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
		ctx.Set("X-Trace-ID", traceID)
		ctx.Set("X-Span-ID", spanID)

		ctxWithDB := context.WithValue(ctx, "DB", db.WithContext(newCtx))
		ctx.Request = ctx.Request.WithContext(ctxWithDB)

		ctx.Next()
	}
}
