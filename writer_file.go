package glog

import (
	"log"
	"fmt"
	"os"
)

func FileLogger(level Level, file *os.File) LogWriter {
	loggers := map[Level]*log.Logger{
		ERROR: log.New(file, fmt.Sprintf("[%s] ", ERROR.String()), DATE_TIME_FILE),
		WARN:  log.New(file, fmt.Sprintf("[%s] ", WARN.String()), DATE_TIME_FILE),
		INFO:  log.New(file, fmt.Sprintf("[%s] ", INFO.String()), DATE_TIME),
		DEBUG: log.New(file, fmt.Sprintf("[%s] ", DEBUG.String()), DATE_TIME),
	}

	return WriterCloserLogger(level, file, loggers)
}
