package repository

import (
	"fmt"
	"go-api-template/configuration"
	"go-api-template/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	db *gorm.DB
}

func NewConnection(cfg *configuration.New) *Connection {
	psqlURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	db, err := gorm.Open(postgres.Open(psqlURL), &gorm.Config{})
	if err != nil {
		logger.Fatalf("connecting to database failed. %v", err)
	}

	logger.Infof("database connection established")

	return &Connection{
		db: db,
	}
}

func (connection *Connection) Close() {
	sqlDB, err := connection.db.DB()
	if err != nil {
		logger.Fatalf("getting database instance failed. %v", err)
	}

	if err = sqlDB.Close(); err != nil {
		logger.Fatalf("closing database connection failed. %v", err)
	}
}
