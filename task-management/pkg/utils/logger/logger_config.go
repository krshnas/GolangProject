package logger

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Zlogger *zap.Logger // Raw zap.Logger instance
var Logger logr.Logger  // Logr.Logger for structured logging

// InitializeLogger initializes the logger with default configuration.
func InitializeLogger() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	Zlogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	Logger = zapr.NewLogger(Zlogger)
	defer Zlogger.Sync()
}

// SetLoggerName dynamically adds a logger name field for identification.
func SetLoggerName(name string) logr.Logger {
	return Logger.WithName(name)
}
