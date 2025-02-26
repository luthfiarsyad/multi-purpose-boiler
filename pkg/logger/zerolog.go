package logger

import (
	"base-app/config"
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// logKey defines keys for context values
type logKey string

const (
	TransactionIDKey logKey = "transaction_id"
	StartTimeKey     logKey = "start_time"
)

// Logger instance
var Logger zerolog.Logger

// Init initializes the logger with JSON format and sets log level
func Init() {
	Logger = zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger()

	zerolog.TimeFieldFormat = "2006-01-02 15:04:05 MST"

	// Set log level from config
	switch config.AppConfig.LogLevel {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel) // Default to info
	}
}

// LogEvent defines structured logging event
type LogEvent struct {
	Level         string      `json:"level"`
	HTTPStatus    int         `json:"http_status"`
	Message       string      `json:"message"`
	TransactionID string      `json:"transaction_id,omitempty"`
	LogPoint      string      `json:"log_point,omitempty"`
	Data          interface{} `json:"data,omitempty"`
	IsBackend     bool        `json:"-"`
}

// LogInfo logs information messages using context
func LogInfo(ctx context.Context, event LogEvent) {
	if event.IsBackend {
		return // Skip backend logs at Info level
	}

	transactionID, _ := ctx.Value(TransactionIDKey).(string)
	startTime, _ := ctx.Value(StartTimeKey).(time.Time)
	processTime := time.Since(startTime).Milliseconds()

	Logger.Info().
		Str("transaction_id", transactionID).
		Int("http_status", event.HTTPStatus).
		Interface("data", event.Data).
		Int64("process_time_ms", processTime).
		Msg(event.Message)
}

// LogError logs error messages using context (always logs regardless of IsBackend)
func LogError(ctx context.Context, event LogEvent, err error) {
	transactionID, _ := ctx.Value(TransactionIDKey).(string)
	startTime, _ := ctx.Value(StartTimeKey).(time.Time)
	processTime := time.Since(startTime).Milliseconds()

	Logger.Error().
		Str("transaction_id", transactionID).
		Int("http_status", event.HTTPStatus).
		Interface("data", event.Data).
		Int64("process_time_ms", processTime).
		Err(err).
		Msg(event.Message)
}

// LogInfoNoCtx logs information messages without context
func LogInfoNoCtx(event LogEvent) {
	if event.IsBackend {
		return // Skip backend logs at Info level
	}

	Logger.Info().
		Str("transaction_id", event.TransactionID).
		Str("log_point", event.LogPoint).
		Int("http_status", event.HTTPStatus).
		Interface("data", event.Data).
		Msg(event.Message)
}

// LogErrorNoCtx logs error messages without context (always logs regardless of IsBackend)
func LogErrorNoCtx(event LogEvent, err error) {
	Logger.Error().
		Str("transaction_id", event.TransactionID).
		Str("log_point", event.LogPoint).
		Int("http_status", event.HTTPStatus).
		Interface("data", event.Data).
		Err(err).
		Msg(event.Message)
}
