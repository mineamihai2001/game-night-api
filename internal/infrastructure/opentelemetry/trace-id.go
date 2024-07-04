package opentelemetry

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

func GetTraceId(ctx context.Context) string {
	spanCtx := trace.SpanFromContext(ctx)

	traceId := spanCtx.SpanContext().TraceID().String()

	return traceId
}
