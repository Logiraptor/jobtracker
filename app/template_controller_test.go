package app

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestTemplateController(t *testing.T) {
	tmpl := template.New("root")
	tmpl.Parse(`Root: {{.AppContext.Port}}`)
	tmpl.New("port.html").Parse(`Port: {{.AppContext.Port}}`)
	var port = rand.Int()
	var controller = TemplateController{
		Template: tmpl,
		AppContext: Context{
			Port:   port,
			Logger: NilLogger{},
		},
	}
	var recorder = httptest.NewRecorder()
	var req = mustNewRequest(t, "GET", "/", nil)
	controller.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, recorder.Body.String(), fmt.Sprintf("Root: %d", port))

	recorder = httptest.NewRecorder()
	req = mustNewRequest(t, "GET", "/nosuchthing", nil)
	controller.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)

	recorder = httptest.NewRecorder()
	req = mustNewRequest(t, "GET", "/port", nil)
	controller.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, recorder.Body.String(), fmt.Sprintf("Port: %d", port))

	recorder = httptest.NewRecorder()
	req = mustNewRequest(t, "GET", "/port.html", nil)
	controller.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, recorder.Body.String(), fmt.Sprintf("Port: %d", port))
}