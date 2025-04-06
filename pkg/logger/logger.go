package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

// responseWriter is a custom writer that captures the response body.
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (response *responseWriter) Write(body []byte) (int, error) {
	response.body.Write(body)

	return response.ResponseWriter.Write(body)
}

// prettyJSONEncoder is a custom zapcore.Encoder that pretty-prints JSON logs.
type prettyJSONEncoder struct {
	zapcore.Encoder
}

func (e *prettyJSONEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	buf, err := e.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, buf.Bytes(), "", "  "); err != nil {
		return nil, err
	}

	b := buffer.NewPool().Get()
	if _, err := b.Write(prettyJSON.Bytes()); err != nil {
		return nil, err
	}

	return b, nil
}

func InitLogger() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	jsonEncoder := zapcore.NewJSONEncoder(config)
	prettyEncoder := &prettyJSONEncoder{Encoder: jsonEncoder}
	core := zapcore.NewCore(prettyEncoder, zapcore.AddSync(os.Stdout), zap.InfoLevel)
	rawLogger := zap.New(core)
	logger = rawLogger.Sugar()
}

func RequestsLogHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Read the request body
		var requestBody []byte
		if context.Request.Body != nil {
			requestBody, _ = io.ReadAll(context.Request.Body)
			context.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// Sanitize the request body
		sanitizedRequestBody := sanitizeAndPrettyPrintBody(requestBody)

		// Create a custom response writer
		customWriter := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: context.Writer}
		context.Writer = customWriter

		// Process request
		context.Next()

		// Calculate latency
		latency := time.Since(startTime)

		// Get request details
		clientIP := context.ClientIP()
		method := context.Request.Method
		path := context.Request.URL.Path
		queryParams := context.Request.URL.Query()
		userAgent := context.Request.UserAgent()
		referer := context.Request.Referer()
		requestID := context.GetHeader("X-Request-ID")
		host := context.Request.Host

		// Get response details
		statusCode := context.Writer.Status()
		bodySize := context.Writer.Size()
		responseBody := customWriter.body.String()

		// Sanitize the response body
		sanitizedResponseBody := sanitizeAndPrettyPrintBody([]byte(responseBody))

		// Decode sanitized and pretty-printed bodies back to map
		var requestBodyMap, responseBodyMap map[string]interface{}
		if err := json.Unmarshal([]byte(sanitizedRequestBody), &requestBodyMap); err != nil {
			requestBodyMap = nil
		}

		if err := json.Unmarshal([]byte(sanitizedResponseBody), &responseBodyMap); err != nil {
			responseBodyMap = nil
		}

		// Log details in JSON format
		logger.Infow("Request details",
			"client_ip", clientIP,
			"method", method,
			"status_code", statusCode,
			"body_size", bodySize,
			"request_body", requestBodyMap,
			"response_body", responseBodyMap,
			"query_params", queryParams,
			"path", path,
			"user_agent", userAgent,
			"referer", referer,
			"request_id", requestID,
			"host", host,
			"latency_ms", latency.Milliseconds(),
		)
	}
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatalf(format string, arges ...interface{}) {
	logger.Fatalf(format, arges...)
}

// sanitizeAndPrettyPrintBody removes or masks sensitive fields from the body and pretty-prints it.
func sanitizeAndPrettyPrintBody(body []byte) string {
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return string(body)
	}

	// List of sensitive fields to remove or mask
	sensitiveFields := []string{"secret", "token", "password", "access_token", "client_secret", "client_id"}
	for _, field := range sensitiveFields {
		if _, exists := data[field]; exists {
			data[field] = "****"
		}
	}

	sanitizedBody, err := json.Marshal(data)
	if err != nil {
		return string(body)
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, sanitizedBody, "", "  "); err != nil {
		return string(sanitizedBody)
	}

	return prettyJSON.String()
}
