package glog

import (
	"log"
	"fmt"
	"os"
)

type FileLogWriter struct {
	*LoggerLogWriter

	File *os.File
}

func FileWriter(level Level, file *os.File) *FileLogWriter {
	loggers := map[Level]*log.Logger{
		ERROR: log.New(file, fmt.Sprintf("[%s] ", ERROR.String()), DATE_TIME_FILE),
		WARN:  log.New(file, fmt.Sprintf("[%s] ", WARN.String()), DATE_TIME_FILE),
		INFO:  log.New(file, fmt.Sprintf("[%s] ", INFO.String()), DATE_TIME),
		DEBUG: log.New(file, fmt.Sprintf("[%s] ", DEBUG.String()), DATE_TIME),
	}

	return &FileLogWriter{File: file, LoggerLogWriter: LoggerWriter(level, loggers)}
}

func (w *FileLogWriter) Write(b []byte) (n int, err error) {
	return w.File.Write(b)
}

func (w *FileLogWriter) Close() error {
	return w.File.Close()
}
