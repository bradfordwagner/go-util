package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

var DefaultConfig = Config{
	Level: "info",
}

// Config is the configuration for the logger.
type Config struct {
	Level string
}

// Init initializes the logger with the given configuration.
func Init(conf Config) {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	level, err := zapcore.ParseLevel(conf.Level)
	if err != nil {
		level = zapcore.InfoLevel
	}
	c := zap.Config{
		EncoderConfig:    encoderCfg,
		Encoding:         "console",
		ErrorOutputPaths: []string{"stderr"},
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stderr"},
	}
	l, _ := c.Build()
	logger = l.Sugar()
}

func Log() *zap.SugaredLogger {
	if logger == nil {
		Init(DefaultConfig)
	}
	return logger
}
