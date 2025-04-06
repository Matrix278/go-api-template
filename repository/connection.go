package repository

import (
	"context"
	"fmt"
	"go-api-template/configuration"
	"go-api-template/pkg/logger"

	"github.com/XSAM/otelsql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // driver for PostgreSQL
	"go.opentelemetry.io/otel"

	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

type Connection struct {
	db *sqlx.DB
}

func NewConnection(cfg *configuration.Env) *Connection {
	psqlURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
		cfg.PostgresSSLMode,
	)

	// Create an instrumented sql.DB
	sqlDB, err := otelsql.Open("postgres", psqlURL,
		otelsql.WithAttributes(
			semconv.DBSystemPostgreSQL,
			semconv.DBNameKey.String(cfg.PostgresDB),
		),
		otelsql.WithTracerProvider(otel.GetTracerProvider()),
		otelsql.WithSpanOptions(otelsql.SpanOptions{
			Ping:     true,
			RowsNext: true,
		}),
	)
	if err != nil {
		logger.Fatalf("failed to create database connection: %v", err)
	}

	// Wrap sql.DB with sqlx
	db := sqlx.NewDb(sqlDB, "postgres")

	// Test the connection with context
	if err = db.PingContext(context.Background()); err != nil {
		logger.Fatalf("connecting to database failed: %v", err)
	}

	logger.Infof("database connection established")

	return &Connection{
		db: db,
	}
}

func (connection *Connection) Close() {
	if err := connection.db.Close(); err != nil {
		logger.Fatalf("closing database connection failed. %v", err)
	}
}
