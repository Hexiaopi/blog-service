package logger

type Logger interface {
	Debug(msg string, fields ...Field)
	Debugf(format string, args ...interface{})
	Info(msg string, fields ...Field)
	Infof(format string, args ...interface{})
	Warn(msg string, fields ...Field)
	Warnf(format string, args ...interface{})
	Error(msg string, fields ...Field)
	Errorf(format string, args ...interface{})
	Panic(msg string, fields ...Field)
	Panicf(format string, args ...interface{})
	Fatal(msg string, fields ...Field)
	Fatalf(format string, args ...interface{})
}

func (l *ZapLogger) Debug(msg string, fields ...Field) {
	l.zapLogger.Debug(msg, fields...)
}

func (l *ZapLogger) Debugf(format string, args ...interface{}) {
	l.zapLogger.Sugar().Debugf(format, args...)
}

func (l *ZapLogger) Info(msg string, fields ...Field) {
	l.zapLogger.Info(msg, fields...)
}

func (l *ZapLogger) Infof(format string, args ...interface{}) {
	l.zapLogger.Sugar().Infof(format, args...)
}

func (l *ZapLogger) Warn(msg string, fields ...Field) {
	l.zapLogger.Warn(msg, fields...)
}

func (l *ZapLogger) Warnf(format string, args ...interface{}) {
	l.zapLogger.Sugar().Warnf(format, args...)
}

func (l *ZapLogger) Error(msg string, fields ...Field) {
	l.zapLogger.Error(msg, fields...)
}

func (l *ZapLogger) Errorf(format string, args ...interface{}) {
	l.zapLogger.Sugar().Errorf(format, args...)
}

func (l *ZapLogger) Panic(msg string, fields ...Field) {
	l.zapLogger.Panic(msg, fields...)
}

func (l *ZapLogger) Panicf(format string, args ...interface{}) {
	l.zapLogger.Sugar().Panicf(format, args...)
}

func (l *ZapLogger) Fatal(msg string, fields ...Field) {
	l.zapLogger.Fatal(msg, fields...)
}

func (l *ZapLogger) Fatalf(format string, args ...interface{}) {
	l.zapLogger.Sugar().Fatalf(format, args...)
}
