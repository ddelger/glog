package glog

type MultiLogWriter struct {
	Writers []LogWriter
}

func MultiLogger(writers ...LogWriter) LogWriter {
	return &MultiLogWriter{Writers: writers}
}

func (w *MultiLogWriter) Log(level Level, calldepth int, format *string, args ...interface{}) {
	for _, writer := range w.Writers {
		if writer.LogLevel() >= level {
			writer.Log(level, calldepth+1, format, args...)
		}
	}
}

func (w *MultiLogWriter) Write(b []byte) (n int, err error) {
	for _, writer := range w.Writers {
		n, err = writer.Write(b)
	}
	return
}

func (w *MultiLogWriter) Close() (err error) {
	for _, writer := range w.Writers {
		err = writer.Close()
	}
	return
}

func (w *MultiLogWriter) LogLevel() Level {
	l := ERROR
	for _, writer := range w.Writers {
		if writer.LogLevel() > l {
			l = writer.LogLevel()
		}
	}
	return l
}
