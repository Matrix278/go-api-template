package repository

import (
	"fmt"
	"go-api-template/configuration"
	"go-api-template/pkg/logger"
	"net"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // driver for PostgreSQL
)

type Connection struct {
	db *sqlx.DB
}

func NewConnection(cfg *configuration.Env) *Connection {
	hostPort := net.JoinHostPort(cfg.PostgresHost, cfg.PostgresPort)
	psqlURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		hostPort,
		cfg.PostgresDB,
		cfg.PostgresSSLMode,
	)

	database, err := sqlx.Open("postgres", psqlURL)
	if err != nil {
		logger.Fatalf("connecting to database failed. %v", err)
	}

	if err = database.Ping(); err != nil {
		logger.Fatalf("connecting to database failed. %v", err)
	}

	logger.Infof("database connection established")

	return &Connection{
		db: database,
	}
}

func (connection *Connection) Close() {
	if err := connection.db.Close(); err != nil {
		logger.Fatalf("closing database connection failed. %v", err)
	}
}
