package telemetry

import (
	ctx "context"
	"errors"
	"go-api-template/pkg/logger"
)

type ShutdownHandler struct {
	funcs []func(ctx.Context) error
}

// NewShutdownHandler creates a new shutdown handler
func NewShutdownHandler() *ShutdownHandler {
	return &ShutdownHandler{
		funcs: make([]func(ctx.Context) error, 0),
	}
}

// AddFunction adds a shutdown function to the handler
func (handle *ShutdownHandler) AddFunction(function func(ctx.Context) error) {
	handle.funcs = append(handle.funcs, function)
}

// Shutdown executes all registered shutdown functions
func (handle *ShutdownHandler) Shutdown(ctx ctx.Context) error {
	var err error
	for _, functions := range handle.funcs {
		if ferr := functions(ctx); ferr != nil {
			err = errors.Join(err, ferr)
			logger.Errorf("shutdown function failed: %v", ferr)
		}
	}

	handle.funcs = nil
	return err
}
