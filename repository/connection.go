package repository

import (
	"fmt"
	"go-api-template/configuration"
	"go-api-template/pkg/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // driver for PostgreSQL
)

type Connection struct {
	db *sqlx.DB
}

func NewConnection(cfg *configuration.Config) *Connection {
	psqlURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
		cfg.PostgresSSLMode,
	)

	db, err := sqlx.Open("postgres", psqlURL)
	if err != nil {
		logger.Fatalf("connecting to database failed. %v", err)
	}

	if err = db.Ping(); err != nil {
		logger.Fatalf("connecting to database failed. %v", err)
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
