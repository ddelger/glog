package glog

import (
	"fmt"
	"log"
	"os"
)

func FileLogger(level Level, file *os.File) LogWriter {
	loggers := map[Level]*log.Logger{
		LevelError: log.New(file, fmt.Sprintf("[%s] ", LevelError.String()), FlagsDateTimeFile),
		LevelWarn:  log.New(file, fmt.Sprintf("[%s] ", LevelWarn.String()), FlagsDateTimeFile),
		LevelInfo:  log.New(file, fmt.Sprintf("[%s] ", LevelInfo.String()), FlagsDateTime),
		LevelDebug: log.New(file, fmt.Sprintf("[%s] ", LevelDebug.String()), FlagsDateTime),
	}

	return WriterCloserLogger(level, file, loggers)
}
