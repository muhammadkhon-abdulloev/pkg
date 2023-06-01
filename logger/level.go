package logger

import "go.uber.org/zap/zapcore"

var logLevels = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(level string) zapcore.Level {
	lvl, ok := logLevels[level]
	if !ok {
		return zapcore.DebugLevel
	}
	return lvl
}
