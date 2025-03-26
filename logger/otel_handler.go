package logger

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"
)

type TracingHandler struct {
	handler slog.Handler
}

func NewTracingHandler(handler slog.Handler) *TracingHandler {
	return &TracingHandler{
		handler: handler,
	}
}

const xtraceid = "x-trace-id"
const xspanid = "x-span-id"

func (h *TracingHandler) Handle(ctx context.Context, r slog.Record) error {
	// 从context中获取span信息
	span := trace.SpanFromContext(ctx)
	if span != nil {
		r.AddAttrs(
			slog.String(xtraceid, span.SpanContext().TraceID().String()),
			slog.String(xspanid, span.SpanContext().SpanID().String()),
		)
	}
	return h.handler.Handle(ctx, r)
}

func (h *TracingHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewTracingHandler(h.handler.WithAttrs(attrs))
}

func (h *TracingHandler) WithGroup(name string) slog.Handler {
	return NewTracingHandler(h.handler.WithGroup(name))
}

func (h *TracingHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}
