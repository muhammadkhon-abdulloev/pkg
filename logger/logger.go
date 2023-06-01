package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

type Logger struct {
	sugaredLogger *zap.SugaredLogger
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) SetLogger(logger *zap.Logger) {
	l.sugaredLogger = logger.Sugar()
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
	l.sugaredLogger.Info(args...)
}

func (l *Logger) Infof(template string, args ...any) {
	l.sugaredLogger.Infof(template, args...)
}

func (l *Logger) Warn(args ...any) {
	l.sugaredLogger.Warn(args...)
}

func (l *Logger) Warnf(template string, args ...any) {
	l.sugaredLogger.Warnf(template, args...)
}

func (l *Logger) Error(args ...any) {
	l.sugaredLogger.Error(args...)
}

func (l *Logger) Errorf(template string, args ...any) {
	l.sugaredLogger.Errorf(template, args...)
}

func (l *Logger) Debug(args ...any) {
	l.sugaredLogger.Debug(args...)
}

func (l *Logger) Debugf(template string, args ...any) {
	l.sugaredLogger.Debugf(template, args...)
}

func (l *Logger) Fatal(args ...any) {
	l.sugaredLogger.Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...any) {
	l.sugaredLogger.Fatalf(template, args...)
}

func (l *Logger) Panic(args ...any) {
	l.sugaredLogger.Panic(args...)
}

func (l *Logger) Panicf(template string, args ...any) {
	l.sugaredLogger.Panicf(template, args...)
}

func (l *Logger) DPanic(args ...any) {
	l.sugaredLogger.DPanic(args...)
}

func (l *Logger) DPanicf(template string, args ...any) {
	l.sugaredLogger.DPanicf(template, args...)
}
