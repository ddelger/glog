package glog

import (
	"log"
	"fmt"
	"io"
)

const (
	DATE_TIME      = log.Ldate | log.Ltime
	DATE_TIME_FILE = log.Ldate | log.Ltime | log.Lshortfile
)

type LoggerLogWriter struct {
	*LeveledLogWriter

	Loggers map[Level]*log.Logger
}

func LoggerWriter(level Level, loggers map[Level]*log.Logger) *LoggerLogWriter {
	return &LoggerLogWriter{LeveledLogWriter: &LeveledLogWriter{Level: level}, Loggers: loggers}
}

func (w *LoggerLogWriter) WithLeveledLogger(level Level, logger *log.Logger) *LoggerLogWriter {
	w.Loggers[level] = logger
	return w
}

func (w *LoggerLogWriter) Log(level Level, calldepth int, format *string, args ...interface{}) {
	var s string
	if format == nil {
		s = fmt.Sprintln(args...)
	} else {
		s = fmt.Sprintf(*format, args...)
	}

	if l, ok := w.Loggers[level]; ok {
		l.Output(calldepth+1, s)
	}
}

type WriterLogWriter struct {
	*LoggerLogWriter

	Writer io.Writer
}

func WriterLogger(level Level, w io.Writer, loggers map[Level]*log.Logger) LogWriter {
	return &WriterLogWriter{Writer: w, LoggerLogWriter: LoggerWriter(level, loggers)}
}

func (w *WriterLogWriter) Write(b []byte) (n int, err error) {
	return w.Writer.Write(b)
}

func (w *WriterLogWriter) Close() error {
	return nil
}

type WriterCloserLogWriter struct {
	LogWriter

	Closer io.Closer
}

func WriterCloserLogger(level Level, wc io.WriteCloser, loggers map[Level]*log.Logger) LogWriter {
	return &WriterCloserLogWriter{Closer: wc, LogWriter: WriterLogger(level, wc, loggers)}
}

func (w *WriterCloserLogWriter) Close() error {
	return w.Closer.Close()
}
