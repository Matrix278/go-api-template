package telemetry

import (
	"context"
	"crypto/tls"
	"fmt"
	"go-api-template/configuration"
	"go-api-template/pkg/logger"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"google.golang.org/grpc/credentials"
)

func InitTracer(cfg *configuration.Telemetry) func() {
	ctx := context.Background()
	shutdown, err := newOtelCollector(ctx, cfg)
	if err != nil {
		logger.Fatalf("failed to setup OpenTelemetry: %v", err)
	}

	return func() {
		logger.Infof("Shutting down OpenTelemetry")
		if err := shutdown(context.Background()); err != nil {
			logger.Errorf("failed to shutdown OpenTelemetry: %v", err)
		}
	}
}

func newOtelCollector(ctx context.Context, cfg *configuration.Telemetry) (func(context.Context) error, error) {
	// Cleanup functions which need to be executed when the OpenTelemetry SDK is shutting down.
	// When shutdown is called (typically when application is terminating),
	// it executes all these cleanup functions in order and combines their errors
	shutdownHandler := NewShutdownHandler()

	// Setup OpenTelemetry SDK
	resources, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithSchemaURL(semconv.SchemaURL),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(cfg.ServiceName),
			semconv.ServiceVersionKey.String(cfg.ServiceVersion),
			semconv.DeploymentEnvironmentKey.String(cfg.Environment),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Otel SDK resource: %w", err)
	}

	// Configure OTEL Exporter
	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(cfg.Endpoint),
		otlptracegrpc.WithHeaders(cfg.Headers),
		otlptracegrpc.WithTLSCredentials(credentials.NewTLS(&tls.Config{
			MinVersion: tls.VersionTLS12,
		})),
		otlptracegrpc.WithTimeout(30 * time.Second),
		otlptracegrpc.WithCompressor(cfg.Compression),
	}

	client := otlptracegrpc.NewClient(opts...)
	traceExporter, err := otlptrace.New(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	// Create BatchSpanProcessor
	batchSpanProcessor := sdktrace.NewBatchSpanProcessor(traceExporter,
		sdktrace.WithMaxQueueSize(cfg.QueueSize),
		sdktrace.WithBatchTimeout(2*time.Second),
		sdktrace.WithMaxExportBatchSize(cfg.MaxExportBatchSize),
	)

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(resources),
		sdktrace.WithSpanProcessor(batchSpanProcessor),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)
	shutdownHandler.AddFunction(tracerProvider.Shutdown)

	// Set global TracerProvider
	otel.SetTracerProvider(tracerProvider)

	// Set global Propagator
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	// Enable runtime metrics
	if err := runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second)); err != nil {
		return nil, fmt.Errorf("failed to start runtime metrics: %w", err)
	}

	logger.Infof("Successfully initialized OpenTelemetry")
	return shutdownHandler.Shutdown, nil
}
