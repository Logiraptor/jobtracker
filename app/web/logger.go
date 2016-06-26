package web

type Logger interface {
	Log(format string, args ...interface{})
}

type NilLogger struct{}

func (n NilLogger) Log(string, ...interface{}) {}
