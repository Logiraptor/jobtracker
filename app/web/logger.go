package web

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
)

type httpRecorder struct {
	code   int
	length int
	http.ResponseWriter
}

func (h *httpRecorder) WriteHeader(code int) {
	h.code = code
	h.ResponseWriter.WriteHeader(code)
}

func (h *httpRecorder) Write(buf []byte) (int, error) {
	if h.code == 0 {
		h.code = 200
	}
	h.length += len(buf)
	return h.ResponseWriter.Write(buf)
}

func LoggerMiddleware(logger *logrus.Logger, handler http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		recorder := &httpRecorder{ResponseWriter: rw}
		handler.ServeHTTP(recorder, req)

		logger.WithFields(logrus.Fields{
			"path":  req.URL.Path,
			"code":  recorder.code,
			"bytes": recorder.length,
			"time":  time.Since(start),
		}).Print("request")
	}
}
