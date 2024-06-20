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

var levelConversion = map[Level]zapcore.Level{
	DebugLevel:  zap.DebugLevel,
	InfoLevel:   zap.InfoLevel,
	WarnLevel:   zap.WarnLevel,
	ErrorLevel:  zap.ErrorLevel,
	DPanicLevel: zap.DPanicLevel,
	PanicLevel:  zap.PanicLevel,
	FatalLevel:  zap.FatalLevel,
}

func init() {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	c := zap.Config{
		EncoderConfig:    encoderCfg,
		Encoding:         "console",
		ErrorOutputPaths: []string{"stderr"},
		Level:            zap.NewAtomicLevelAt(levelConversion[LevelConfig]),
		OutputPaths:      []string{"stderr"},
	}
	l, _ := c.Build()
	logger = l.Sugar()
}

var LevelConfig Level = InfoLevel

func Log() *zap.SugaredLogger {
	return logger
}
