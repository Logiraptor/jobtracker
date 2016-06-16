package app

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSessionsController(t *testing.T) {
	var recorder = httptest.NewRecorder()
	var request, _ = http.NewRequest("POST", "/login", strings.NewReader(url.Values{
		"email":    {"email@example.com"},
		"password": {"password"},
	}.Encode()))
	var controller = SessionsController{}
	controller.Create(recorder, request)
	assert.Equal(t, http.StatusFound, recorder.Code)
	assert.NotEmpty(t, recorder.HeaderMap.Get("Set-Cookie"), "Returns a user session cookie")
}
