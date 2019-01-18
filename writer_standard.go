package glog

import (
	"fmt"
	"log"
	"os"

	"github.com/mgutz/ansi"
)

func StandardLogger(level Level) LogWriter {
	loggers := map[Level]*log.Logger{
		LevelError: log.New(os.Stderr, fmt.Sprintf("%s[%s]%s ", ansi.Red, LevelError.String(), ansi.Reset), FlagsDateTimeFile),
		LevelWarn:  log.New(os.Stdout, fmt.Sprintf("%s[%s]%s ", ansi.Yellow, LevelWarn.String(), ansi.Reset), FlagsDateTimeFile),
		LevelInfo:  log.New(os.Stdout, fmt.Sprintf("[%s] ", LevelInfo.String()), FlagsDateTime),
		LevelDebug: log.New(os.Stdout, fmt.Sprintf("[%s] ", LevelDebug.String()), FlagsDateTime),
	}

	return WriterLogger(level, os.Stdout, loggers)
}
