package glog

import (
	"log"
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

type StandardLogWriter struct {
	*LoggerLogWriter
}

func StandardWriter(level Level) *StandardLogWriter {
	loggers := map[Level]*log.Logger{
		ERROR: log.New(os.Stderr, fmt.Sprintf("%s[%s]%s ", ansi.Red, ERROR.String(), ansi.Reset), DATE_TIME_FILE),
		WARN:  log.New(os.Stdout, fmt.Sprintf("%s[%s]%s ", ansi.Yellow, WARN.String(), ansi.Reset), DATE_TIME_FILE),
		INFO:  log.New(os.Stdout, fmt.Sprintf("[%s] ", INFO.String()), DATE_TIME),
		DEBUG: log.New(os.Stdout, fmt.Sprintf("[%s] ", DEBUG.String()), DATE_TIME),
	}

	return &StandardLogWriter{LoggerLogWriter: LoggerWriter(level, loggers)}
}

func (w *StandardLogWriter) Write(b []byte) (n int, err error) {
	return os.Stdout.Write(b)
}

func (w *StandardLogWriter) Close() error {
	return nil
}
