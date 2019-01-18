package glog

import (
	"fmt"
	"os"
)

var logger = &Logger{Writer: StandardLogger(LevelInfo)}

type Logger struct {
	Writer LogWriter
}

func Loggers(writers ...LogWriter) {
	logger.Writer = MultiLogger(writers...)
}

func (l *Logger) log(level Level, format *string, args ...interface{}) {
	if l.Writer.LogLevel() >= level {
		l.Writer.Log(level, 4, format, args...)
	}
}

func (l *Logger) Fatal(args ...interface{}) {
	l.log(LevelError, nil, args...)
	os.Exit(1)
}
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.log(LevelError, &format, args...)
	os.Exit(1)
}

func (l *Logger) Panic(args ...interface{}) {
	l.log(LevelError, nil, args...)
	panic(fmt.Sprintln(args...))
}
func (l *Logger) Panicf(format string, args ...interface{}) {
	l.log(LevelError, &format, args...)
	panic(fmt.Sprintf(format, args...))
}

func (l *Logger) Error(args ...interface{}) {
	l.log(LevelError, nil, args...)
}
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.log(LevelError, &format, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.log(LevelWarn, nil, args...)
}
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.log(LevelWarn, &format, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.log(LevelInfo, nil, args...)
}
func (l *Logger) Infof(format string, args ...interface{}) {
	l.log(LevelInfo, &format, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.log(LevelDebug, nil, args...)
}
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.log(LevelDebug, &format, args...)
}

func Fatal(args ...interface{})                 { logger.Fatal(args...) }
func Fatalf(format string, args ...interface{}) { logger.Fatalf(format, args...) }
func Panic(args ...interface{})                 { logger.Panic(args...) }
func Panicf(format string, args ...interface{}) { logger.Panicf(format, args...) }
func Errorf(format string, args ...interface{}) { logger.Errorf(format, args...) }
func Error(args ...interface{})                 { logger.Error(args...) }
func Warn(args ...interface{})                  { logger.Warn(args...) }
func Warnf(format string, args ...interface{})  { logger.Warnf(format, args...) }
func Info(args ...interface{})                  { logger.Info(args...) }
func Infof(format string, args ...interface{})  { logger.Infof(format, args...) }
func Debug(args ...interface{})                 { logger.Debug(args...) }
func Debugf(format string, args ...interface{}) { logger.Debugf(format, args...) }
