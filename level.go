package glog

type Level int

const (
	ERROR Level = iota
	WARN
	INFO
	DEBUG
	NONE
)

var LEVELS = map[Level]string{
	ERROR: "ERROR",
	WARN:  "WARN",
	INFO:  "INFO",
	DEBUG: "DEBUG",
}

func (l Level) String() string {
	return LEVELS[l]
}
