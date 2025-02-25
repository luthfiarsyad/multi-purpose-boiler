// cmd/main.go
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

	// Log informasi server saat start
	logger.LogInfo(logger.LogEvent{
		Level:         "INFO",
		HTTPStatus:    http.StatusOK,
		Message:       "Starting HTTP server",
		TransactionID: "tx-server-start",
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
	logger.LogInfo(logger.LogEvent{
		Level:         "INFO",
		HTTPStatus:    http.StatusOK,
		Message:       "Server listening on " + serv.Addr,
		TransactionID: "tx-server-run",
	})

	if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.LogError(logger.LogEvent{
			Level:         "ERROR",
			HTTPStatus:    http.StatusInternalServerError,
			Message:       "Server failed to start",
			TransactionID: "tx-server-run",
		}, err)
		errChan <- err
	}
}

func waitForShutdown(serv *http.Server, errChan <-chan error) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	logger.LogInfo(logger.LogEvent{
		Level:         "INFO",
		HTTPStatus:    http.StatusOK,
		Message:       "Waiting for shutdown signal or server error",
		TransactionID: "tx-server-wait",
	})

	select {
	case sig := <-sigChan:
		logger.LogInfo(logger.LogEvent{
			Level:         "INFO",
			HTTPStatus:    http.StatusOK,
			Message:       "Received signal: " + sig.String(),
			TransactionID: "tx-server-shutdown",
		})
		shutdownGracefully(serv)
	case err := <-errChan:
		logger.LogError(logger.LogEvent{
			Level:         "ERROR",
			HTTPStatus:    http.StatusInternalServerError,
			Message:       "Server encountered an error",
			TransactionID: "tx-server-wait",
		}, err)
		panic(err)
	}
}

func shutdownGracefully(serv *http.Server) {
	logger.LogInfo(logger.LogEvent{
		Level:         "INFO",
		HTTPStatus:    http.StatusOK,
		Message:       "Initiating graceful shutdown",
		TransactionID: "tx-server-shutdown",
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := serv.Shutdown(ctx); err != nil {
		logger.LogError(logger.LogEvent{
			Level:         "ERROR",
			HTTPStatus:    http.StatusInternalServerError,
			Message:       "Shutdown failed",
			TransactionID: "tx-server-shutdown",
		}, err)
		panic(err)
	}

	logger.LogInfo(logger.LogEvent{
		Level:         "INFO",
		HTTPStatus:    http.StatusOK,
		Message:       "Server gracefully stopped",
		TransactionID: "tx-server-shutdown",
	})
}
