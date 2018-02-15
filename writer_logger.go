package glog

import (
	"log"
	"fmt"
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

func (w *LoggerLogWriter) Log(level Level, calldepth int, format *string, args ...interface{}) {
	var s string
	if format == nil {
		s = fmt.Sprintln(args...)
	} else {
		s = fmt.Sprintf(*format, args...)
	}

	w.Loggers[level].Output(calldepth+1, s)
}
