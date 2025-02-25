package logger

import (
	"os"

	"github.com/rs/zerolog"
)

// Logger adalah instance global zerolog.Logger
var Logger zerolog.Logger

// Init menginisialisasi logger dengan konfigurasi JSON
func Init() {
	Logger = zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger()

	zerolog.TimeFieldFormat = "2006-01-02 15:04:05 MST"
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

// LogEvent adalah struktur untuk logging terstruktur
type LogEvent struct {
	Level         string
	HTTPStatus    int
	Message       string
	TransactionID string
	Data          interface{}
}

// LogInfo mencetak log dengan level INFO
func LogInfo(event LogEvent) {
	Logger.Info().
		Str("level", event.Level).
		Int("http_status", event.HTTPStatus).
		Str("transaction_id", event.TransactionID).
		Interface("data", event.Data).
		Msg(event.Message)
}

// LogError mencetak log dengan level ERROR
func LogError(event LogEvent, err error) {
	Logger.Error().
		Str("level", event.Level).
		Int("http_status", event.HTTPStatus).
		Str("transaction_id", event.TransactionID).
		Interface("data", event.Data).
		Err(err).
		Msg(event.Message)
}
