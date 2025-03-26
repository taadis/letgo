package middleware

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type TraceMiddleware struct {
	tracer     trace.Tracer
	propagator propagation.TextMapPropagator
}

func NewTraceMiddleware(tracer trace.Tracer) *TraceMiddleware {
	return &TraceMiddleware{
		tracer:     tracer,
		propagator: propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	}
}

func (m *TraceMiddleware) extractContext(r *http.Request) context.Context {
	ctx := r.Context()

	// 先尝试从标准 traceparent 提取
	ctx = m.propagator.Extract(ctx, propagation.HeaderCarrier(r.Header))

	// 如果没有标准 traceparent，则尝试使用自定义 x-trace-id
	span := trace.SpanFromContext(ctx)
	if !span.SpanContext().HasTraceID() {
		// custom x-trace-id
		if traceID := r.Header.Get("x-trace-id"); traceID != "" {
			if tid, err := trace.TraceIDFromHex(traceID); err == nil {
				spanCtx := trace.NewSpanContext(trace.SpanContextConfig{
					TraceID:    tid,
					SpanID:     trace.SpanID{},
					TraceFlags: trace.FlagsSampled,
				})
				ctx = trace.ContextWithSpanContext(ctx, spanCtx)
			}
		}
	}

	return ctx
}

func (m *TraceMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := m.extractContext(r)
		ctx, span := m.tracer.Start(ctx, "http-request")
		defer span.End()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *TraceMiddleware) HandlerFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := m.extractContext(r)
		ctx, span := m.tracer.Start(ctx, "http-request")
		defer span.End()

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
