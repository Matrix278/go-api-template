package configuration

import (
	"strings"

	"github.com/spf13/viper"
)

type Telemetry struct {
	ServiceName        string
	ServiceVersion     string
	Environment        string
	Endpoint           string
	Headers            map[string]string
	Insecure           bool
	QueueSize          int
	MaxExportBatchSize int
	Compression        string
}

// TelemetryNew creates a new Telemetry configuration from environment
func TelemetryNew() *Telemetry {
	viper.SetDefault("OTEL_SERVICE_NAME", "go-api-template")
	viper.SetDefault("OTEL_SERVICE_VERSION", "1.0.0")
	viper.SetDefault("OTEL_EXPORTER_OTLP_ENDPOINT", "localhost:4317")
	viper.SetDefault("OTEL_INSECURE", true)
	viper.SetDefault("OTEL_EXPORTER_OTLP_HEADERS", "")
	viper.SetDefault("OTEL_EXPORTER_OTLP_QUEUE_SIZE", 4096)
	viper.SetDefault("OTEL_EXPORTER_OTLP_MAX_EXPORT_BATCH_SIZE", 512)
	viper.SetDefault("OTEL_EXPORTER_OTLP_COMPRESSION", "gzip")

	headers := make(map[string]string)
	if headerStr := viper.GetString("OTEL_EXPORTER_OTLP_HEADERS"); headerStr != "" {
		for _, header := range strings.Split(headerStr, ",") {
			parts := strings.SplitN(header, "=", 2)
			if len(parts) == 2 {
				headers[parts[0]] = parts[1]
			}
		}
	}

	return &Telemetry{
		ServiceName:        viper.GetString("OTEL_SERVICE_NAME"),
		ServiceVersion:     viper.GetString("OTEL_SERVICE_VERSION"),
		Environment:        viper.GetString("APP_ENV"),
		Endpoint:           viper.GetString("OTEL_EXPORTER_OTLP_ENDPOINT"),
		Headers:            headers,
		Insecure:           viper.GetBool("OTEL_INSECURE"),
		QueueSize:          viper.GetInt("OTEL_EXPORTER_OTLP_QUEUE_SIZE"),
		MaxExportBatchSize: viper.GetInt("OTEL_EXPORTER_OTLP_MAX_EXPORT_BATCH_SIZE"),
		Compression:        viper.GetString("OTEL_EXPORTER_OTLP_COMPRESSION"),
	}
}
