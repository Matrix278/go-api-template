package main

import (
	"context"
	"errors"
	"go-api-template/configuration"
	"go-api-template/controller"
	"go-api-template/middleware"
	"go-api-template/model"
	"go-api-template/pkg/logger"
	"go-api-template/repository"
	"go-api-template/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "gorm.io/driver/postgres"
)

//	@title			Go API
//	@version		1.0
//	@description	GO API documentation

// @securityDefinitions.basic	BasicAuth.
func main() {
	// Initialize the logger
	logger.InitLogger()

	// Load the configuration
	cfg, err := configuration.Load()
	if err != nil {
		logger.Errorf("configuration loading failed: %v", err)
	}

	// Initialize validation
	if err := model.InitValidation(); err != nil {
		log.Fatalf("failed to initialize validation: %v", err)
	}

	// Initialize the database connection
	dbConnection := repository.NewConnection(cfg)
	defer dbConnection.Close()

	// Initialize repositories
	repository := repository.NewRepositories(dbConnection)

	// Initialize services
	services := service.NewServices(cfg, repository)

	// Initialize the controllers
	controllers := controller.NewControllers(services)

	// Initialize router
	router, err := middleware.NewRouter(cfg, controllers)
	if err != nil {
		logger.Errorf("initializing router failed: %v", err)
	}

	srv := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
		// set timeout due CWE-400 - Potential Slowloris Attack
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		logger.Infof("Listening server on port %s...", cfg.AppPort)

		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("Server failed to start: %v", err)
		}
	}()

	gracefulShutdown(srv)
}

func gracefulShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logger.Infof("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Infof("Server exiting")
}
