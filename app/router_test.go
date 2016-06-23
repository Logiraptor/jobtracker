package app

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/manveru/faker"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func mustNewRequest(t *testing.T, method string, path string, bodyValues url.Values) *http.Request {
	var body io.Reader = strings.NewReader(bodyValues.Encode())
	if bodyValues == nil {
		body = nil
	}
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err.Error())
		return nil
	}
	if bodyValues != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	var buf = &bytes.Buffer{}
	req.Write(buf)
	req, err = http.ReadRequest(bufio.NewReader(buf))
	if err != nil {
		t.Fatalf("Failed to read request: %s", err.Error())
		return nil
	}
	return req
}

type TestLogger struct {
	logs []string
}

func (t *TestLogger) Log(fmtString string, args ...interface{}) {
	t.logs = append(t.logs, fmt.Sprintf(fmtString, args...))
}

func TestPather(t *testing.T) {
	var fake, _ = faker.New("en")
	var definedPathName = fake.Characters(10)
	var undefinedPathName = fake.Characters(10)

	routes := mux.NewRouter()
	routes.NewRoute().Path("/test_path").Name(definedPathName)
	var logger = &TestLogger{}
	p := NewPather(logger, routes)
	assert.Equal(t, p.Path(definedPathName), "/test_path")
	assert.Equal(t, p.Path(undefinedPathName), "/")
	assert.Len(t, logger.logs, 1)
	assert.Contains(t, logger.logs[0], undefinedPathName)
}
