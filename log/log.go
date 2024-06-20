package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

type Level string

const (
	DebugLevel  Level = "debug"
	InfoLevel   Level = "info"
	WarnLevel   Level = "warn"
	ErrorLevel  Level = "error"
	DPanicLevel Level = "dpanic"
	PanicLevel  Level = "panic"
	FatalLevel  Level = "fatal"
)

var levelConversion = map[Level]zap.Level{
	DebugLevel:  zapcore.DebugLevel,
	InfoLevel:   zapcore.InfoLevel,
	WarnLevel:   zapcore.WarnLevel,
	ErrorLevel:  zapcore.ErrorLevel,
	DPanicLevel: zapcore.DPanicLevel,
	PanicLevel:  zapcore.PanicLevel,
	FatalLevel:  zapcore.FatalLevel,
}

type Config struct {
	Level zapcore.Level
}

// SetLevel sets the log level for the logger.
func (c *Config) SetLevel(level string) {
	ok, l := levelConversion[Level(level)]
	if !ok {
		l = InfoLevel
	}
	c.Level = l
}

func Init(c Config) {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	c := zap.Config{
		EncoderConfig:    encoderCfg,
		Encoding:         "console",
		ErrorOutputPaths: []string{"stderr"},
		Level:            zap.NewAtomicLevelAt(c.Level),
		OutputPaths:      []string{"stderr"},
	}
	l, _ := c.Build()
	logger = l.Sugar()
}

var LevelConfig Level = InfoLevel

func Log() *zap.SugaredLogger {
	return logger
}
