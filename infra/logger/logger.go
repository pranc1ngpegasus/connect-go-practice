package logger

import (
	"context"
	"fmt"

	domain "github.com/Pranc1ngPegasus/connect-go-practice/domain/logger"
	"github.com/google/wire"
	"github.com/samber/lo"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ domain.Logger = (*Logger)(nil)

var NewLoggerSet = wire.NewSet(
	wire.Bind(new(domain.Logger), new(*Logger)),
	NewLogger,
)

type Logger struct {
	logger *otelzap.Logger
}

func NewLogger() (*Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	l, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	return &Logger{
		logger: otelzap.New(l),
	}, nil
}

func encoderConfig() zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()
	cfg.LevelKey = "severity"
	cfg.EncodeLevel = EncodeLevel
	cfg.TimeKey = "time"
	cfg.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	return cfg
}

var logLevelSeverity = map[zapcore.Level]string{
	zapcore.DebugLevel:  "DEBUG",
	zapcore.InfoLevel:   "INFO",
	zapcore.WarnLevel:   "WARNING",
	zapcore.ErrorLevel:  "ERROR",
	zapcore.DPanicLevel: "CRITICAL",
	zapcore.PanicLevel:  "ALERT",
	zapcore.FatalLevel:  "EMERGENCY",
}

func EncodeLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(logLevelSeverity[l])
}

func (l *Logger) Field(key string, iface interface{}) domain.Field {
	return domain.Field{
		Key:       key,
		Interface: iface,
	}
}

func (l *Logger) field(field domain.Field) zap.Field {
	switch i := field.Interface.(type) {
	case error:
		return zap.Error(i)
	case string:
		return zap.String(field.Key, i)
	case int:
		return zap.Int(field.Key, i)
	case bool:
		return zap.Bool(field.Key, i)
	default:
		return zap.Any(field.Key, i)
	}
}

func (l *Logger) Debug(ctx context.Context, message string, fields ...domain.Field) {
	zapfields := lo.Map(fields, func(field domain.Field, _ int) zap.Field {
		return l.field(field)
	})

	l.logger.Ctx(ctx).Debug(message, zapfields...)
}

func (l *Logger) Info(ctx context.Context, message string, fields ...domain.Field) {
	zapfields := lo.Map(fields, func(field domain.Field, _ int) zap.Field {
		return l.field(field)
	})

	l.logger.Ctx(ctx).Info(message, zapfields...)
}

func (l *Logger) Error(ctx context.Context, message string, fields ...domain.Field) {
	zapfields := lo.Map(fields, func(field domain.Field, _ int) zap.Field {
		return l.field(field)
	})

	l.logger.Ctx(ctx).Error(message, zapfields...)
}
