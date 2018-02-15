package glog

import "io"

type LogWriter interface {
	io.WriteCloser

	Log(Level, int, *string, ...interface{})
	LogLevel() Level
}
