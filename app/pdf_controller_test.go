package app

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"rsc.io/pdf"

	"github.com/stretchr/testify/assert"
)

func TestPdfController(t *testing.T) {
	var controller = PdfController{
		Logger: NilLogger{},
	}
	var recorder = httptest.NewRecorder()
	var request, _ = http.NewRequest("GET", "/pdf", nil)
	controller.Generate(recorder, request)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "application/pdf", recorder.Header().Get("Content-Type"))
	rd, err := pdf.NewReader(bytes.NewReader(recorder.Body.Bytes()), int64(recorder.Body.Len()))
	assert.Nil(t, err)
	assert.Equal(t, 1, rd.NumPage())
}
