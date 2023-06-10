package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) SetLogger(logger *zap.Logger) {
	l.logger = logger
}

func (l *Logger) InitLogger(c Config) error {
	var (
		encoderCfg zapcore.EncoderConfig
		encoder    zapcore.Encoder
		writer     io.Writer
	)
	if c.ToFile {
		file, err := os.OpenFile(c.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("os.OpenFile: %w", err)
		}
		writer = file
	} else {
		writer = os.Stdout
	}

	wSyncer := zapcore.AddSync(writer)

	if c.Production {
		encoderCfg = zap.NewProductionEncoderConfig()
	} else {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	}

	encoderCfg.LevelKey = LevelKey
	encoderCfg.CallerKey = CallerKey
	encoderCfg.TimeKey = TimeKey
	encoderCfg.NameKey = NameKey
	encoderCfg.MessageKey = MessageKey
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	if c.Decoder == JSONDecoder {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	}

	core := zapcore.NewCore(
		encoder,
		wSyncer,
		zap.NewAtomicLevelAt(getLoggerLevel(c.Level)),
	)

	l.SetLogger(zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(c.AddCallerSkip),
	))

	return nil
}

func (l *Logger) Info(args ...any) {
	l.logger.Sugar().Info(args...)
}

func (l *Logger) Infof(template string, args ...any) {
	l.logger.Sugar().Infof(template, args...)
}

func (l *Logger) Warn(args ...any) {
	l.logger.Sugar().Warn(args...)
}

func (l *Logger) Warnf(template string, args ...any) {
	l.logger.Sugar().Warnf(template, args...)
}

func (l *Logger) Error(args ...any) {
	l.logger.Sugar().Error(args...)
}

func (l *Logger) Errorf(template string, args ...any) {
	l.logger.Sugar().Errorf(template, args...)
}

func (l *Logger) Debug(args ...any) {
	l.logger.Sugar().Debug(args...)
}

func (l *Logger) Debugf(template string, args ...any) {
	l.logger.Sugar().Debugf(template, args...)
}

func (l *Logger) Fatal(args ...any) {
	l.logger.Sugar().Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...any) {
	l.logger.Sugar().Fatalf(template, args...)
}

func (l *Logger) Panic(args ...any) {
	l.logger.Sugar().Panic(args...)
}

func (l *Logger) Panicf(template string, args ...any) {
	l.logger.Sugar().Panicf(template, args...)
}

func (l *Logger) DPanic(args ...any) {
	l.logger.Sugar().DPanic(args...)
}

func (l *Logger) DPanicf(template string, args ...any) {
	l.logger.Sugar().DPanicf(template, args...)
}

func (l *Logger) Log(lvl zapcore.Level, msg string, fields ...zapcore.Field) {
	l.logger.Log(lvl, msg, fields...)
}
