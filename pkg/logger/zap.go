package logger

import (
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZapLogger struct {
	zapLogger *zap.Logger
}

type Config struct {
	FileName  string
	LogLevel  string
	MaxSize   int
	MaxBackup int
	MaxAge    int
	Compress  bool
	Encoding  string
	Env       string
}

func New(conf *Config) Logger {
	lv := conf.LogLevel
	var level zapcore.Level
	switch strings.ToLower(lv) {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	hook := lumberjack.Logger{
		Filename:   conf.FileName,
		MaxSize:    conf.MaxSize,
		MaxBackups: conf.MaxBackup,
		MaxAge:     conf.MaxAge,
		Compress:   conf.Compress,
	}

	var encoder zapcore.Encoder
	if conf.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "Logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
			EncodeTime:     timeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
		})
	} else {
		encoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     timeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook), zapcore.AddSync(os.Stdout)),
		level,
	)
	if conf.Env != "prod" {
		return &ZapLogger{zap.New(core, zap.Development(), zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
	}
	return &ZapLogger{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.InfoLevel))}
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	//enc.AppendString(t.Format("2006-01-02 15:04:05"))
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// // WithValue Adds a field to the specified context
// func (l *Logger) WithValue(ctx context.Context, fields ...zapcore.Field) context.Context {
// 	if c, ok := ctx.(*gin.Context); ok {
// 		ctx = c.Request.Context()
// 		c.Request = c.Request.WithContext(context.WithValue(ctx, ctxLoggerKey, l.WithContext(ctx).With(fields...)))
// 		return c
// 	}
// 	return context.WithValue(ctx, ctxLoggerKey, l.WithContext(ctx).With(fields...))
// }

// // WithContext Returns a zap instance from the specified context
// func (l *Logger) WithContext(ctx context.Context) *Logger {
// 	if c, ok := ctx.(*gin.Context); ok {
// 		ctx = c.Request.Context()
// 	}
// 	zl := ctx.Value(ctxLoggerKey)
// 	ctxLogger, ok := zl.(*zap.Logger)
// 	if ok {
// 		return &Logger{ctxLogger}
// 	}
// 	return l
// }
