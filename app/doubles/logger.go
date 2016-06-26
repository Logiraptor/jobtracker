package doubles

import "fmt"

type FakeLogger struct {
	Logs []string
}

func (t *FakeLogger) Log(fmtString string, args ...interface{}) {
	t.Logs = append(t.Logs, fmt.Sprintf(fmtString, args...))
}
