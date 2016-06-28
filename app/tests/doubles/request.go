package doubles

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func NewRequest(t *testing.T, method string, path string, bodyValues url.Values) *http.Request {
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
