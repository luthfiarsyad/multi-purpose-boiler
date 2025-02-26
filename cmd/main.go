package main

import (
	"base-app/cmd/server"
	"base-app/pkg/logger"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Inisialisasi logger sebelum server
	logger.Init()

	// Dapatkan instance server dari server.go
	serv := server.GetServer()
	errChan := make(chan error, 1)

	// Log informasi server saat start tanpa context
	logger.LogInfoNoCtx(logger.LogEvent{
		Level:      "INFO",
		HTTPStatus: http.StatusOK,
		Message:    "Starting HTTP server",
		LogPoint:   "server-start",
		Data: map[string]interface{}{
			"addr":          serv.Addr,
			"read_timeout":  serv.ReadTimeout.String(),
			"write_timeout": serv.WriteTimeout.String(),
		},
	})

	go runServer(serv, errChan)
	waitForShutdown(serv, errChan)
}

func runServer(serv *http.Server, errChan chan<- error) {
	logger.LogInfoNoCtx(logger.LogEvent{
		Level:      "INFO",
		HTTPStatus: http.StatusOK,
		Message:    "Server listening on " + serv.Addr,
		LogPoint:   "server-running",
	})

	if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.LogErrorNoCtx(logger.LogEvent{
			Level:      "ERROR",
			HTTPStatus: http.StatusInternalServerError,
			Message:    "Server failed to start",
			LogPoint:   "server-error",
		}, err)
		errChan <- err
	}
}

func waitForShutdown(serv *http.Server, errChan <-chan error) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	logger.LogInfoNoCtx(logger.LogEvent{
		Level:      "INFO",
		HTTPStatus: http.StatusOK,
		Message:    "Waiting for shutdown signal or server error",
		LogPoint:   "server-wait",
	})

	select {
	case sig := <-sigChan:
		logger.LogInfoNoCtx(logger.LogEvent{
			Level:      "INFO",
			HTTPStatus: http.StatusOK,
			Message:    "Received signal: " + sig.String(),
			LogPoint:   "server-shutdown",
		})
		shutdownGracefully(serv)
	case err := <-errChan:
		logger.LogErrorNoCtx(logger.LogEvent{
			Level:      "ERROR",
			HTTPStatus: http.StatusInternalServerError,
			Message:    "Server encountered an error",
			LogPoint:   "server-error",
		}, err)
		panic(err)
	}
}

func shutdownGracefully(serv *http.Server) {
	logger.LogInfoNoCtx(logger.LogEvent{
		Level:      "INFO",
		HTTPStatus: http.StatusOK,
		Message:    "Initiating graceful shutdown",
		LogPoint:   "shutdown-start",
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := serv.Shutdown(ctx); err != nil {
		logger.LogErrorNoCtx(logger.LogEvent{
			Level:      "ERROR",
			HTTPStatus: http.StatusInternalServerError,
			Message:    "Shutdown failed",
			LogPoint:   "shutdown-failed",
		}, err)
		panic(err)
	}

	logger.LogInfoNoCtx(logger.LogEvent{
		Level:      "INFO",
		HTTPStatus: http.StatusOK,
		Message:    "Server gracefully stopped",
		LogPoint:   "shutdown-success",
	})
}
