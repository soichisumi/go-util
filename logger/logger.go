package logger

import (
	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"sync"
)

const calldepth = 1


var (
	l    *zap.Logger
	once sync.Once
)

func init() {
	once.Do(func() {
		l = new()
	})
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(msg string, fields ...zap.Field) {
	l.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(msg string, fields ...zap.Field) {
	l.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(msg string, fields ...zap.Field) {
	l.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(msg string, fields ...zap.Field) {
	l.Error(msg, fields...)
}

// DPanic logs a message at DPanicLevel. The message includes any fields
// passed at the log site, as well as any fields accumulated on the logger.
//
// If the logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
func DPanic(msg string, fields ...zap.Field) {
	l.DPanic(msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func Panic(msg string, fields ...zap.Field) {
	l.Panic(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func Fatal(msg string, fields ...zap.Field) {
	l.Fatal(msg, fields...)
}

// Logger ...
func Logger() *zap.Logger {
	return l
}

func SetOptions(opts ...zap.Option) {
	l = l.WithOptions(opts...)
}

func new() *zap.Logger {
	var cfg zap.Config
	switch os.Getenv("ENV") {
	case "bench":
		cfg = benchConfig()
	default:
		cfg = defaultConfig()
	}

	logger, err := cfg.Build(zapdriver.WrapCore(), zap.AddCallerSkip(calldepth))
	if err != nil {
		log.Fatalln(err)
	}

	return logger
}

func defaultConfig() zap.Config {
	var cfg = zapdriver.NewDevelopmentConfig()
	cfg.InitialFields = map[string]interface{}{"env": os.Getenv("ENV")}
	cfg.DisableStacktrace = true
	return cfg
}

func benchConfig() zap.Config {
	var cfg = zapdriver.NewProductionConfig()
	cfg.InitialFields = map[string]interface{}{"env": os.Getenv("ENV")}
	cfg.DisableStacktrace = true
	cfg.Level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
	return cfg
}