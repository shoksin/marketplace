package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

func Init(serviceName string) {
	cfg := zap.NewProductionConfig()

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.MessageKey = "message"

	cfg.InitialFields = map[string]interface{}{
		"service": serviceName,
		"env":     os.Getenv("APP_ENV"), //production, staging
		"version": os.Getenv("APP_VERSION"),
	}

	baseLogger, err := cfg.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		panic("Failed to build logger: " + err.Error())
	}
	Logger = baseLogger
}

func CloseSync() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}
