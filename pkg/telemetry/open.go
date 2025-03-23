package telemetry

import (
	"context"
	"crypto/tls"
	"errors"
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
	shutdown, err := setupOTelSDK(ctx, cfg)
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

func setupOTelSDK(ctx context.Context, cfg *configuration.Telemetry) (func(context.Context) error, error) {
	var shutdownFuncs []func(context.Context) error

	// Create shutdown function
	shutdown := func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	// Create resource
	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithSchemaURL(semconv.SchemaURL),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(cfg.ServiceName),
			semconv.ServiceVersionKey.String(cfg.ServiceVersion),
			semconv.DeploymentEnvironmentKey.String(cfg.Environment),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// Configure OTLP exporter
	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(cfg.Endpoint),
		otlptracegrpc.WithHeaders(cfg.Headers),
		otlptracegrpc.WithTLSCredentials(credentials.NewTLS(&tls.Config{
			MinVersion: tls.VersionTLS12,
		})),
		otlptracegrpc.WithTimeout(30 * time.Second),
		otlptracegrpc.WithCompressor("gzip"),
	}

	client := otlptracegrpc.NewClient(opts...)
	traceExporter, err := otlptrace.New(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	// Create BatchSpanProcessor
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter,
		sdktrace.WithMaxQueueSize(2048),
		sdktrace.WithBatchTimeout(2*time.Second),
		sdktrace.WithMaxExportBatchSize(512),
	)

	// Create TracerProvider
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)

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
	return shutdown, nil
}
