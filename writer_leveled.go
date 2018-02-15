package glog

type LeveledLogWriter struct {
	Level Level
}

func (w *LeveledLogWriter) LogLevel() Level {
	return w.Level
}
