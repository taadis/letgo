package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/taadis/letgo/logger"
	"github.com/taadis/letgo/middleware"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

func initLogger() {
	handler := logger.NewTracingHandler(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
	)
	slog.SetDefault(slog.New(handler))
}

func initTracer() (*trace.TracerProvider, error) {
	// exporter, err := stdouttrace.New()
	// if err != nil {
	// 	return nil, err
	// }
	tp := trace.NewTracerProvider(
	//trace.WithSampler(trace.AlwaysSample()),
	//trace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}

func main() {
	initLogger()
	tp, err := initTracer()
	if err != nil {
		slog.Error("failed to init tracer", "error", err)
		return
	}
	defer tp.Shutdown(context.Background())

	tracer := tp.Tracer("main-tracer")

	// 创建一个根 span 用于服务器生命周期
	ctx, span := tracer.Start(context.Background(), "server-lifecycle")
	defer span.End()

	traceMiddleware := middleware.NewTraceMiddleware(tracer)

	http.HandleFunc("/", traceMiddleware.HandlerFunc(Home))
	http.HandleFunc("/hello", traceMiddleware.HandlerFunc(Api))
	http.HandleFunc("/none", None)
	slog.InfoContext(ctx, "starting server")
	http.ListenAndServe(":8080", nil)
}
