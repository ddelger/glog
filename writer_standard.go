package glog

import (
	"log"
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

func StandardLogger(level Level) LogWriter {
	loggers := map[Level]*log.Logger{
		ERROR: log.New(os.Stderr, fmt.Sprintf("%s[%s]%s ", ansi.Red, ERROR.String(), ansi.Reset), DATE_TIME_FILE),
		WARN:  log.New(os.Stdout, fmt.Sprintf("%s[%s]%s ", ansi.Yellow, WARN.String(), ansi.Reset), DATE_TIME_FILE),
		INFO:  log.New(os.Stdout, fmt.Sprintf("[%s] ", INFO.String()), DATE_TIME),
		DEBUG: log.New(os.Stdout, fmt.Sprintf("[%s] ", DEBUG.String()), DATE_TIME),
	}

	return WriterLogger(level, os.Stdout, loggers)
}
