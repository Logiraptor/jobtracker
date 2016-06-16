package app

import (
	"net/http"
	"path/filepath"
	"strconv"
)

type Context struct {
	Logger  Logger
	AppRoot string
	Port    int
}

func Start(ctx Context) error {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		http.ServeFile(rw, req, filepath.Join(ctx.AppRoot+"public/index.html"))
	})

	var pdfController = PdfController{
		Logger: ctx.Logger,
	}

	http.HandleFunc("/pdf", pdfController.Generate)

	ctx.Logger.Log("App started on port: %d", ctx.Port)

	go http.ListenAndServe(":"+strconv.Itoa(ctx.Port), nil)

	return nil
}

type NilLogger struct{}

func (n NilLogger) Log(string, ...interface{}) {}
